## Web Skeleton

Web Skeleton is a simple http/net server set up with routing, template rendering with the only external library being Mux for routing.

Implementing Web Skeleton
Using Web Skeleton is simple. First use go get to install the latest version of the library.

	$ go get github.com/spectre013/webskeleton


Once you have the source navigate to webskeleton directory and run. 

	$ go run main.go

This will start the server at [http://localhost:3000](http://localhost:3000) 

The handlers are all defined in handlers.go and rendering remplates is done my calling the render function.


	Render(w, "index", "Test")

the render function has the abillity to render multiple templates. If you wanted to add a sidebar you would add 

	<? template "sidebar" . ?>

to the index.html file. In sidebar.html file you would need to define the content.

	<? define "sidebar" ?>
		<!-- ... some html here ... -->
	<? end ?>

then would need to add the template to the render call.

	Render(w, "index,sidebar", "Test")
