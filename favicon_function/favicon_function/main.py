from fastapi import FastAPI
from . import extract_favicon

app = FastAPI()


@app.get("/")
async def root():
    return {"message": "Hello World"}


@app.get("/favicon/{fqdn}", status_code=200)
async def get_favicon(fqdn):
    extract_favicon.extract_favicon(fqdn)
    return
