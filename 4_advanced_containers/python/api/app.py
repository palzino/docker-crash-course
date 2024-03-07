from flask import Flask
import os
def create_app():
    app = Flask(__name__)
    @app.route('/get')
    def get_name():
        # Fetch the environment variable "NAME"; return a default if it's not set
        name = os.getenv('NAME', 'No name set')
        return {'name': name}
    app.run(debug=True)
    return app
app = create_app() 
if __name__ == '__main__':
    app.run(debug=True)
