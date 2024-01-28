FROM arigaio/atlas:latest

WORKDIR app/
COPY backend/atlas.hcl ./atlas.hcl
COPY backend/migrations migrations