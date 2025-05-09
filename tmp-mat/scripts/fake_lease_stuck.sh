set -x
docker exec -it spqr_qdb_0_1 etcdctl get "/2p_commits/" --prefix
lease=$(docker exec -it spqr_qdb_0_1 etcdctl lease grant 30 | awk '{ print $2 }')
docker exec -it spqr_qdb_0_1 etcdctl put "/2p_commits/lease/123" 123 --lease="$lease"
docker exec -it spqr_qdb_0_1 etcdctl put "/2p_commits/123" "Init"
docker exec -it spqr_qdb_0_1 etcdctl get "/2p_commits/" --prefix
