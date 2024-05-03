from typing import Callable

from fastapi import APIRouter, FastAPI, Request, Response
from fastapi.routing import APIRoute

import logging
import json

from datetime import datetime

app = FastAPI()

logging.basicConfig(
   filename='logs/server.log', 
   level=logging.INFO, 
   format='%(message)s'
)

class LoggedRoute(APIRoute):
    def get_route_handler(self) -> Callable:
        original_route_handler = super().get_route_handler()

        async def custom_route_handler(request: Request) -> Response:
            response: Response = await original_route_handler(request)

            log_data = {
                'time': datetime.now().isoformat(),
                'response': response.body.decode('utf-8'),
            }
            logging.info(json.dumps(log_data))

            return response
        
        return custom_route_handler


router = APIRouter(route_class=LoggedRoute)

@router.get("/api/")
def read_item(userid: int):
    return {"userid": userid}

app.include_router(router)