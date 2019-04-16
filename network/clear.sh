echo STOP AND DELETE THE DOCKER CONTAINERS
docker ps -aq | xargs -n 1 docker stop
docker ps -aq | xargs -n 1 docker rm -v

echo DELETE THE OLD DOCKER VOLUMES
docker volume prune -f

echo DELETE OLD DOCKER NETWORKS
docker network prune -f

echo Remove all stopped containers
docker container prune -f

echo Removing All Unused Objects
docker system prune -f
docker system prune --volumes -f

echo DELETE SCRIPT-CREATED FILES
rm -rf channel-artifacts/*.block channel-artifacts/*.tx crypto-config
