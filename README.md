# Project's Title

The Ultimate Timer

# Project Description

Have you ever been irritated to set your timer by second so that you can include the additional time from you start the timer to do your stuff?

Or, have you ever wanted to save your timer as preset so that you can use them everyday work.

The Ultimate Timer is here to set you free from those annoying moments.

You can save your custom timer presets in The Ultimate Timer.

# Contents

The project is created by Typescript and Golang.

The frontend is organized quite simply as being made upon the Next.js framework, whereas the backend is structured by Golang with DDD.

Hence, Redis is utilized in the backend to fetch the requested resource immediately

If Redis does not have the requested resource, it will then look for the resource in the Postgres.

if the resourse was found in the Postgres, it will cache the data concurrently.

Although as written above, the backend is made according to the DDD architecture, I realized that there is so much room for improvement after I learned more in-depth about DDD.

So, honestly, I am a little bit nervous about publishing this project, but still, sharing what you have learned is what the developers' culture is, right?

# How to Install and Run the Project

git clone this project.

`git clone https://github.com/SDTakeuchi/ultimate_timer.git`

Then, docker-compose up the project at the directory where docker-compose.yaml is.

`docker-compose up --build`

# How to Use the Project

### CAUTION
#### I am afraid that the frontend is work in progress because I wanted to learn specifically about the backend knowledge. So, there are plenty of unnecessary buttons and inputs here and there...


You can use the project on your favorite browser.

After you successfully docker-compose up the project, open your browser and visit `http://localhost/timer`. 

I am afraid that the first visit to the page will frequently timeout, but please try several times to see the page.

At `http://localhost/timer`, you can see the list of your saved timer presets.

If you want to create your first timer preset, you can push the `CREATE A NEW PRESET` button.

And on the page, You can create your timer preset by following instructions.

1. Fill in "Timer Preset Name"
1. Click the "Add a unit" button
1. Put "1" in the "order" field
1. Put desired timer duration by second in the "duration" field
1. Finally, "SUBMIT"

Back in the `http://localhost/timer`, you can click the play button for each timer preset to go to the detail page.

On the next page, you can push either the "PLAY" button or the "RESTART" button to start the timer.

API is also accessible. The endpoints are listed below.

GET `http://localhost/api/presets/` : retrieve a list of presets

GET `http://localhost/api/presets/{id}` : retrieve a preset specified by ID

POST `http://localhost/api/presets/` : create a new preset

PUT `http://localhost/api/presets/{id}` : update a preset specified by ID

DELETE `http://localhost/api/presets/{id}` : delete a preset specified by ID
