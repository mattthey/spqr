/*__spqr__engine_v2:true,__spqr__commit_strategy:2pc*/ begin;
/*__spqr__engine_v2:true,__spqr__commit_strategy:2pc*/ update orders set customer_id = customer_id + 1 where id = 2;
/*__spqr__engine_v2:true,__spqr__commit_strategy:2pc*/ update orders set customer_id = customer_id + 1 where id = 1001;
/*__spqr__engine_v2:true,__spqr__commit_strategy:2pc*/ commit;