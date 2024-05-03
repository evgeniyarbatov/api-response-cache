from fastapi import Request, Response
from fastapi.routing import APIRoute

from typing import Callable
from datetime import datetime

import logging

logging.basicConfig(
   filename='logs/logs.csv', 
   level=logging.INFO,
   format="%(message)s"
)

class LoggedRoute(APIRoute):
    def get_route_handler(self) -> Callable:
        original_route_handler = super().get_route_handler()

        async def custom_route_handler(request: Request) -> Response:
            response: Response = await original_route_handler(request)
            log_data = [
                datetime.now().isoformat(),
                response.body.decode('utf-8'),
            ]
            logging.info(log_data)
            return response
        
        return custom_route_handler