// service worker for chixm applications.


// install event runs at the first time you accessed the web site.
self.addEventListener('install', function () {
    debugger;
    console.log("Start Install event");
});

// fetch event runs when the user changed url location.
self.addEventListener('fetch', function (e) {
    debugger;
    e.waitUntil(
        console.log("Start Fetch event")
    );
});

// activate is an event when service worker file was updated and replaced.
self.addEventListener('activate', function (event) {
    console.log("Start Activate event");
});