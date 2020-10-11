import web
import json

class App(web.application):
    def run(self, port=8080, *middleware):
        func = self.wsgifunc(*middleware)
        return web.httpserver.runsimple(func, ('0.0.0.0', port))

class Webhook:
    def POST(self):
        # If the POST has a payload
        if data:
            data = json.loads(web.data())
            # Return the raw data we received in the POST payload
            return json.dumps({"response": f'{data}'})
        # If the POST doesn't have a payload
        else:
            return json.dumps({"response": "Empty POST Request Received"})

    def GET(self):
        # Return a default response for on a GET request
        return json.dumps({"response": "No GETs Here Thanks"})

if __name__ == '__main__':
    urls = (
        '/', 'Webhook'
    )

    app = App(urls, globals())
    app.run()
