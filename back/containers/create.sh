mkdir data
docker run -v /full/path/of/data:/data -e ZINC_DATA_PATH="/data" -p 4080:4080 \
    -e ZINC_FIRST_ADMIN_USER=admin -e ZINC_FIRST_ADMIN_PASSWORD=admin \
    --name zincsearch public.ecr.aws/zinclabs/zincsearch:latest