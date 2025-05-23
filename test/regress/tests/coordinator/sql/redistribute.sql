REGISTER ROUTER r1 ADDRESS regress_router:7000;

CREATE DISTRIBUTION ds1 COLUMN TYPES integer;
CREATE KEY RANGE krid1 FROM 1 ROUTE TO sh1 FOR DISTRIBUTION ds1;

SHOW key_ranges;

REDISTRIBUTE KEY RANGE krid1 TO sh1 BATCH SIZE 100;

SHOW key_ranges;

REDISTRIBUTE KEY RANGE krid1 TO sh2 BATCH SIZE 100;

SHOW key_ranges;

REDISTRIBUTE KEY RANGE krid1 TO sh1 BATCH SIZE -1;

REDISTRIBUTE KEY RANGE krid1 TO sh2;

SHOW key_ranges;

DROP DISTRIBUTION ALL CASCADE;
DROP KEY RANGE ALL;
UNREGISTER ROUTER ALL;
