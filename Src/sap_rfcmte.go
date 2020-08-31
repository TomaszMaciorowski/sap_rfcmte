package main

import (
        "net/http"
        "time"
        "fmt"
        "strconv"
        "github.com/sap/gorfc/gorfc"
        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promauto"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

func abapSystem() gorfc.ConnectionParameters {
        return gorfc.ConnectionParameters{
                "user":   "user_rfc",
                "passwd": "password",
                "ashost": "ip",
                "sysnr":  "00",
                "client": "100",
                "lang":   "EN",
        }
}

func GiveSapRFC(mtsysid string, mtmcname string, mtnumrange string, mtuid string, mtclass string , mtindex string , extrindex string) float64 {

        params := map[string]interface{}{
                "MTSYSID"       : mtsysid,
                "MTMCNAME"      : mtmcname,
                "MTNUMRANGE"    : mtnumrange,
                "MTUID"         : mtuid,
                "MTCLASS"       : mtclass,
                "MTINDEX"       : mtindex,
                "EXTINDEX"      : extrindex,
                }
        c, err := gorfc.ConnectionFromParams(abapSystem())
        if err != nil {

		}
        fmt.Println("Connected:", c.Alive())

		r, e := c.Call("ZPROMETHEUS_READ", params)
        if e != nil {

		}
        lastperval := r["LASTPERVAL"]
        fmt.Println(r["CUSGRPNAME"])
        str := fmt.Sprint(lastperval)
        val,err := strconv.ParseFloat(str,64)
        if err != nil {
                }
        c.Close()
        return val
}

func recordMetrics() {
        go func() {
                var val float64 = 0.0;
                for {

                        val = GiveSapRFC("PRD","WAW00SAP0007_PRD_00","010","0000001401","100","0000000267","0000000159")
                        UsersLoggedIn.Set(val)
                        val = GiveSapRFC("PRD","WAW00SAP0007_PRD_00","010","0000001380","100","0000000242","0000000136")
                        ResponseTime.Set(val)
                        val = GiveSapRFC("PRD","WAW00SAP0007_PRD_00","005","0000000005","100","0000000243","0000000137")
                        ResponseTimeDialog.Set(val)
                        val = GiveSapRFC("PRD","WAW00SAP0007_PRD_00","010","0000001391","100","0000000251","0000000145")
                        Utilisation.Set(val)
                        val = GiveSapRFC("PRD","WAW00SAP0007_PRD_00","005","0000000105","100","0000000060","0000000014")
                        UtilisationCpu.Set(val)
                        val = GiveSapRFC("PRD","WAW00SAP0007_PRD_00","010","0000007941","100","0000000433","0000000257")
                        MemoryEsAct.Set(val)
			time.Sleep(60 * time.Second)
                }
        }()
}

var (
        UsersLoggedIn = promauto.NewGauge(prometheus.GaugeOpts{
                Name: "zmysap_UsersLoggedIn",
                Help: "UsersLoggedIn",
        })
        ResponseTime = promauto.NewGauge(prometheus.GaugeOpts{
                Name: "zmysap_ResponseTime",
                Help: "Responsetime czas odpowiedzi",
        })
        ResponseTimeDialog = promauto.NewGauge(prometheus.GaugeOpts{
                Name: "zmysap_ResponseTimeDialog",
                Help: "Responsetimedialog czas odpowiedzi",
        })
        Utilisation = promauto.NewGauge(prometheus.GaugeOpts{
                Name: "zmysap_Utilisation",
                Help: "dialog utilisation %",
        })
        UtilisationCpu = promauto.NewGauge(prometheus.GaugeOpts{
                Name: "zmysap_UtilisationCpu",
                Help: "cpu utilisation %",
        })
        MemoryEsAct = promauto.NewGauge(prometheus.GaugeOpts{
                Name: "zmysap_MemoryEsAct",
                Help: "memory %",
        })
)

func main() {
        recordMetrics()
        http.Handle("/metrics", promhttp.Handler())
        http.ListenAndServe(":2113", nil)
}
