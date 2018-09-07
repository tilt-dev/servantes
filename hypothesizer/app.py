from flask import Flask
import secrets
app = Flask(__name__)

benefits = [
    "Jurassic park is real",
    "You can time travel",
    "You can fall asleep whenever you want",
]

drawbacks = [
    "there's a turtle somewhere in the world, moving toward you at all times, and if it touches you you die",
    "you can only wear one shirt for the rest of your life",
]

@app.route('/')
def make_hypothesis():
    benefit = secrets.choice(benefits)
    drawback = secrets.choice(drawbacks)

    return "{} BUT {}.".format(benefit, drawback)

if __name__ == '__main__':
    print("hello world people!")
    app.run(debug=True, host='0.0.0.0')
