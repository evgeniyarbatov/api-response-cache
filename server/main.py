from fastapi import APIRouter, FastAPI
from logged_route import LoggedRoute

app = FastAPI()
router = APIRouter(route_class=LoggedRoute)

@router.get("/api/")
def read_item(userid: int):
    return {"userid": userid}

app.include_router(router)