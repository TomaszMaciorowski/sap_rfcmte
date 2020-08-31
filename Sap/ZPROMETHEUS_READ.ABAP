function ZPROMETHEUS_READ.
*"----------------------------------------------------------------------
*"*"Lokalny interfejs:
*"  IMPORTING
*"     VALUE(MTSYSID) TYPE  TEXT8 OPTIONAL
*"     VALUE(MTMCNAME) TYPE  TEXT40 OPTIONAL
*"     VALUE(MTNUMRANGE) TYPE  CNUM3 OPTIONAL
*"     VALUE(MTUID) TYPE  TEXT10 OPTIONAL
*"     VALUE(MTCLASS) TYPE  CNUM3 OPTIONAL
*"     VALUE(MTINDEX) TYPE  TEXT10 OPTIONAL
*"     VALUE(EXTINDEX) TYPE  TEXT10 OPTIONAL
*"  EXPORTING
*"     VALUE(CUSGRPNAME) TYPE  STRING
*"     VALUE(VALUNIT) TYPE  STRING
*"     VALUE(LASTPERVAL) TYPE  STRING
*"  EXCEPTIONS
*"      NO_AUTHORITY
*"----------------------------------------------------------------------

DATA:
ld_only_local	 TYPE ALPARAMS-ONLY_LOCAL,
ld_bypass_cache	 TYPE ALFLAG,
it_tid   TYPE STANDARD TABLE OF ALGLOBTID,"TABLES PARAM
wa_tid   LIKE LINE OF it_tid ,
it_perftyperc	 TYPE STANDARD TABLE OF  ALPERFTYPE,"TABLES PARAM
Var2 type string value '',
wa_perftyperc	 LIKE LINE OF it_perftyperc .

wa_tid-MTSYSID = MTSYSID.
wa_tid-MTMCNAME = MTMCNAME.
wa_tid-MTNUMRANGE = MTNUMRANGE.
wa_tid-MTUID = MTUID.
wa_tid-MTCLASS = MTCLASS.
wa_tid-MTINDEX = MTINDEX.
wa_tid-EXTINDEX = EXTINDEX.

append wa_tid to it_tid.

append wa_perftyperc to it_perftyperc.

ld_only_local = ' '.
ld_bypass_cache = ' '.

CALL FUNCTION 'SALC_PERF_READ'

*EXPORTING
* only_local =                 ld_only_local
* bypass_cache =               ld_bypass_cache
  TABLES
    tid =                        it_tid
    perftyperc =                 it_perftyperc
  EXCEPTIONS
    C_CALL_FAILED =              1
    .  "  SALC_PERF_READ

IF SY-SUBRC EQ 0.
  READ table it_perftyperc INTO wa_perftyperc INDEX 1.
  CUSGRPNAME = wa_perftyperc-CUSGRPNAME.
  VALUNIT = wa_perftyperc-VALUNIT.
  LASTPERVAL = wa_perftyperc-LASTPERVAL.
  "All OK
ELSEIF SY-SUBRC EQ 1. "Exception
  "Add code for exception here
   CUSGRPNAME = 'ERROR'.
ENDIF.

endfunction.
