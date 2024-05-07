function getStudent(data){
    const students=[]//to store all student ID
    const allStudents=JSON.parse(data)
    // add student id to students list
    allStudents.forEach(stud=>{
        students.push(stud.stdid)
    });
    // get the sid input field
    var select=document.getElementById("sid")
    // iterate over the students and create a new opton for each ID
    for (var i=0; i<students.length; i++){
        var sid= students[i];
        var option=document.createElement("option")
        option.textContent=sid;
        option.value=sid;
        select.appendChild(option);
    }
}

function getCourses(data){
    const courses=[]
    const allCourses=JSON.parse(data) //The JSON.parse(data) is used to converted the data sendt from the server into javascript objects.
    // it allows us to manipulate the data like without it the data would be a string and we won't be able to loop through the data with 
    // the forEach() loop as done here.
    allCourses.forEach(course=>{
        courses.push(course.cid)
    });

    var option="";
    for(var i=0; i<courses.length; i++){
        option+='<option value="'+courses[i]+'">'+courses[i]+"</option>"
    }
    document.getElementById("cid").innerHTML=option;
}

window.onload=function(){
    // request to get all students
    fetch('/students')
    .then(response=>response.text()) // It reads the entire body of the response and returns it as a string.
    .then(data=>getStudent(data));
    // get all courses
    fetch('/courses')
    .then(response=>response.text())
    .then(data=>getCourses(data));

    fetch('/enrolls')
    .then(response=>response.text())
    .then(data=> getAllEnroll(data));
}

function addEnroll(){
    var _data={
        stdid: parseInt(document.getElementById("sid").value),
        cid: document.getElementById("cid").value,
    }
    var sid=_data.stdid;
    var cid=_data.cid;

    if(isNaN(sid) || cid==""){
        alert("Select valid data")
        return
    }
    fetch('/enroll',{
        method: "POST",
        body: JSON.stringify(_data),
        headers:{"Content-type": "application/json; charset=UTF-8"}
    }).then(response=>{
        if(response.ok){
            fetch('/enroll/'+sid+'/'+cid)
            .then(response=>response.text())
            .then(data=>getEnrolled(data))
        }else{
            throw new Error(response.statusText)
        }
    }).catch(e=>{
        if(e=="Error: Forbidden"){
            alert(e+". Duplicate entry!")
        }
    })
    resetField();
}

function resetField(){
    document.getElementById("sid").value="";
    document.getElementById("cid").value="";
}

function getEnrolled(data){
    const enrolled = JSON.parse(data)
    showTable(enrolled)

}

function showTable(enrolled){
    // Find a <table> element with id="myTable":
    var table=document.getElementById("myTable");
    // Create an empty <tr> element and add it to the last position of the table:
    var row=table.insertRow(table.length);
    // Insert new cells (<td> elements) at the 1st and 2nd position of the "new" <tr> elements:
    var td=[]
    for(i=0; i<table.rows[0].cells.length; i++){
        td[i]=row.insertCell(i);
    }
    td[0].innerHTML= enrolled.stdid;
    td[1].innerHTML= enrolled.cid;
    td[2].innerHTML= enrolled.date.split("T")[0];//show only date ignore time
    td[3].innerHTML= '<input type="button" onclick="deleteEnroll(this)" value="Delete" id="button-1">';
}

function getAllEnroll(data){
    const allenroll=JSON.parse(data)
    allenroll.forEach(enroll=>{
        showTable(enroll)
    });
}

const deleteEnroll=async(r)=>{
    if(confirm('Are you sure you want to delete this?')){
        selectedRow=r.parentElement.parentElement;
        sid= selectedRow.cells[0].innerHTML;
        cid= selectedRow.cells[1].innerHTML;
        fetch('/enroll/'+sid+"/"+cid, {
            method: "DELETE",
            headers:{"Content-type":"application/json; charset=UTF-8"}
        }).then(response=>{
            if(response.ok){
                var rowIndex= selectedRow.rowIndex; //index starts at 0
                if(rowIndex>0){
                    document.getElementById("myTable").deleteRow(rowIndex);
                }
            }
        });
    }
}