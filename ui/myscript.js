window.onload = (event) => {
    console.log('page is fully loaded');
    
    fetch('/projects')
     .then(response => response.json())
        .then(data => {
            console.log(data)
            htmlText = ""
            data.projects.map(project => {
                htmlText += "<li> <b>" + project.Name + "</b> : " + project.Description + "</li>"
            })
            document.getElementById("myprojects").innerHTML = htmlText;
        });

};