To compile worker/controller:
go get -u github.com/go-redis/redis
go get -u github.com/twinj/uuid

to change something:
- make changes in code
- rebuild with increased tag
- publish with increased tag
- goto step 10 and update both deployments with new tag (worker, controller)
- goto step 10 and redeploy.sh
