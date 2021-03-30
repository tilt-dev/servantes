# servantes
A microservice app that demonstrates the power of [Tilt](https://tilt.build).

(Like Miguel de Cervantes, but with servers!)

## Quick start
1. [Install `tilt`](https://docs.tilt.dev/install.html)
1. Start [a local Kubernetes cluster](https://docs.tilt.dev/choosing_clusters.html)
1. Verify you have `m4` (either through XCode Command Line Tools on Mac or `apt install m4` on Linux)
1. `git clone git@github.com:tilt-dev/servantes.git`
1. `cd servantes`
1. `tilt up`

This will deploy several microservice apps in to the default namespace of your current kubernetes cluster. Use the arrow keys to navigate between services in the Tilt HUD. Hit 'b' to open a service in the (B)rowser; the service `fe` is the frontend, and the most interesting to look at.

## Screenshot
![Screenshot of Servantes](/images/frontend.png)

## Demo Workflows
If you're exploring Tilt via Servantes, here are some things you can do to Servantes to see features of Tilt.

### Update Your App
The Snack service is easy to edit. Open the file `snack/main.go` and find the constant strings that are the options for snacks it will recommend. Comment all them out and write "Hello Tilt". Save, and watch Tilt build and update. Then reload Servantes in your web browse and see the new string.

### Pinpoint Problems
Tilt's UX is built to highlight active problems, no matter where they're happening. Here are some ways you can break Servantes and see errors in Tilt.

#### Build Breakage
In `snack/main.go`, type in random characters and save. Tilt will start a build, and when it fails put the error in the HUD.

Fix the error and move on.

#### Startup Error
Kubernetes Pods can get into `CrashLoopBackOff` when they can't startup, e.g. because they can't find a necessary resource file.

### Interactive Onboarding
Tilt's onboarding is nifty because you can do it interactively, as described in our [Tutorial](https://docs.tilt.build/tutorial.html). Instead of writing your config first then running Tilt, Tilt watches your config and updates itself as you configure it.

You can recreate this experience by resetting to a pre-Tilt state and then adding it back in:
1. Stop tilt
2. Run `tilt down` to delete what Tilt has created
3. Open `Tiltfile` in your editor and delete/comment out the entire contents.
4. Run `tilt up` with an empty Tiltfile
5. Add lines back (you'll want to add it back in the order in the original Tiltfile).

You can see Tilt spring to life as it gets more data.

### live\_update
Tilt can [update Kubernetes in seconds, not minutes](https://medium.com/windmill-engineering/how-tilt-updates-kubernetes-in-seconds-not-minutes-28ddffe2d79f) by using `live_update`. Because Servantes is a demo app, most services are small enough that they don't need optimizations. We purposely built our frontend to have a slow build (it links in the Kubernetes client library which can take minutes to build). Servantes uses `live_update` for the frontend to demo Tilt's speed at updating running services.

In `fe/main.go`, change the constant `maxWidgets` to display fewer widgets. Tilt will update `fe` in-place in seconds. Then reload Servantes and you should see your chosen number of services.
