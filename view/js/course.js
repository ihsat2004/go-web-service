// to display all added courses
window.onload = function(){
    fetch("/courses")
    .then(response => response.text())
    .then(data => showCourses(data))
}

//
function addRow(course){
    var table = document.getElementById("myTable");
    var row = table.insertRow(table.length);

    var td= []
    for(i = 0; i<table.rows[0].cells.length; i++){
    td[i] = row.insertCell(i);
    }
    //add course deatil to the new cell
    td[0].innerHTML = course.cid;
    td[1].innerHTML = course.coursename;
    td[2].innerHTML = '<input type = "button" onclick = "deleteCourse(this)"value="delete" id = "button-1">';
    td[3].innerHTML = '<input type = "button" onclick = "updateCourse(this)"value = "edit" id= "button-2">';
}

function showCourses(data){
        const courses = JSON.parse(data)
        courses.forEach(course =>{
            var table = document.getElementById("myTable");
        addRow(course)
    })
}

function showCourse(data){
    const course = JSON.parse(data)
    addRow(course)
}

function getFormData(){
    var data= {
        cid : document.getElementById("cid").value,
        coursename : document.getElementById("cname").value
        }
        return data
}


//adding course
function addCourse(){
    var data = getFormData()
    fetch('/course',{
        method: "POST",
        body: JSON.stringify(data),
        headers: {"content-type": "application/json; charset=UTF-8"}
    }).then(response1 =>{
        var cid = data.cid;
        if(response1.ok){
            fetch('/course/'+cid)
            .then(response2 => response2.text())
            .then(data => showCourse(data))
        }else{
            throw new Error(response1.statusText)
        }
}).catch(e =>{
    alert(e)
})
resetform();
    var cid = data.cid
    if(cid == ""){
        alert("enter a valid Course ID")
        return
    }else if (data.coursename == ""){
        alert("coursename cannot be empty")
    return
    }
    console.log(data) 
   
}

function resetform(){
    document.getElementById("cid").value = "";
    document.getElementById("cname").value = "";
}

//delete button
function deleteCourse(r){
    // this(input) -> td -> tr
    if (confirm('Are you sure you want to DELETE this?')){
    selectedRow = r.parentElement.parentElement;
    cid = selectedRow.cells[0].innerHTML;
    fetch('/course/'+cid, {
    method: "DELETE",
    headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(res => {
        if(res.ok){
            alert("course deleted")
            var rowIndex = selectedRow.rowIndex
            if(rowIndex){
                document.getElementById("myTable").deleteRow(rowIndex)
            }
            selectedRow = null
        }
    })
   }
}

function update(cid){
    //extract new data from the form
    var newData = getFormData()
    fetch("/course/"+cid, {
        method: "PUT",
        //json stringify convert js object into go object
        body: JSON.stringify(newData),
        // headers: {"Content-type":"application/json; charset=UTF-8"}
    }).then(res => {
        if (res.ok){
            //fill in selected row with updated value
            selectedRow.cells[0].innerHTML = newData.courseid;
            selectedRow.cells[1].innerHTML = newData.coursename;

            //set to previous value
            var button = document.getElementById("button-add");
            button.innerHTML = "Add";
            button.setAttribute("onclick", "addCourse()");
            selectedRow = null;

            resetform();

        }else{
            alert("server: Update request error.")
        }
    })
}
//update button
var selectedRow = null
function updateCourse(r){

    //r.parentElement is td or data is stored in td
    selectedRow = r.parentElement.parentElement;

    //filling the form as soon as we click on edit button and update
    document.getElementById("cid").value = selectedRow.cells[0].innerHTML;
    document.getElementById("cname").value = selectedRow.cells[1].innerHTML;  
    
    //
    var btn = document.getElementById("button-add")
    if(btn){
        btn.innerHTML = "update"

cid = selectedRow.cells[0].innerHTML
        btn.setAttribute("onclick","update(cid)")
    }
}