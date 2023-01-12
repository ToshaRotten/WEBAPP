wrong = document.getElementById("wrong");
allInputs = document.getElementsByTagName("input");
wrongs = 0;
var now = new Date();

function showPass() {
    if (allInputs[3].type === "password") {
        allInputs[3].type = "text";
        allInputs[4].type = "text";
    } else {
        allInputs[3].type = "password";
        allInputs[4].type = "password";
    }
}

for (let i = 0 ; i < allInputs.length; i++){
    allInputs.outerHTML = parseInt(allInputs[i].value);
}

function AddWrong(message, inputNumber){
    wrong.style.display = "block";
    wrong.innerHTML = message;
    allInputs[inputNumber].style.borderLeft = "3px solid red";
    wrongs++;
}

function DeleteWrong(inputNumber){
    allInputs[inputNumber].style.border = "1px solid #ced4da";
    wrongs--;
}

function Clear(){
    wrong.style.display = "none";
    for (let i = 0 ; i < allInputs.length; i++){
        DeleteWrong(i);
    }
    wrongs = 0;
}



function Validation() {
    Clear(); // При каждом нажатии отчищать счетчик ошибок ввода и стиль input-ов
    // Имя валидно
    if (! /^[a-zA-Zа-яА-Я]+$/.test(allInputs[0].value)) {
        AddWrong("Имя содержит недопустимые символы" , 0);
    }
    if (allInputs[0].value.length > 30){
        AddWrong("" , 3);
        AddWrong("Имя не должно превышать 30 символов" , 4);
    }

    // Фамилия валидна
    if (! /^[a-zA-Zа-яА-Я]+$/.test(allInputs[1].value)) {
        AddWrong("Фамилия содержит недопустимые символы" , 1);
    }
    if (allInputs[1].value.length > 30){
        AddWrong("" , 3);
        AddWrong("Фамилия не должна превышать 30 символов" , 4);
    }
    

    // Пароль валидный
    if (allInputs[3].value.length < 8){
        AddWrong("" , 3);
        AddWrong("Пароль должен содержать более 8 символов" , 4);
    }
    if (! /^[a-zA-Z0-9а-яА-я!"$%&'()*+,-./:;<=>?@^_`{|}~]+$/.test(allInputs[3].value)){
        AddWrong("" , 3);
        AddWrong("Пароль содержит недопустимые символы" , 4);
    }
    if ( /^[^!"$%&'()*+,-./:;<=>?@^_`{|}~]*$/.test(allInputs[3].value)){
        AddWrong("" , 3);
        AddWrong("Пароль должен содержать хотябы один специальный символ: [^!" +"$/%/&" +"'()*+,-./:;<=>?@^_`{|}~]" , 4);
    }
    if ( /^[^a-zа-я]*$/.test(allInputs[3].value)){
        AddWrong("" , 3);
        AddWrong("Пароль должен содержать хотябы одну прописную букву" , 4);
    }
    if ( /^[^A-ZА-Я]*$/.test(allInputs[3].value)){
        AddWrong("" , 3);
        AddWrong("Пароль должен содержать хотябы одну заглавную букву" , 4);
    }
    if ( /^[^0-9]*$/.test(allInputs[3].value)){
        AddWrong("" , 3);
        AddWrong("Пароль должен содержать хотябы одну цифру" , 4);
    }

    
    // Пароль и подтверждение совпадает
    if (allInputs[3].value != allInputs[4].value){
        AddWrong("" , 3);
        AddWrong("Пароли должны совпадать" , 4);
    }


    // Email валидный
    if (! /^[a-zA-Z0-9@.]+$/.test(allInputs[2].value)) {
        AddWrong("Email содержит недопустимые символы" , 2);
    }
    if (!/^(([^<>()[\].,;:\s@"]+(\.[^<>()[\].,;:\s@"]+)*)|(".+"))@(([^<>()[\].,;:\s@"]+\.)+[^<>()[\].,;:\s@"]{2,})$/.test(allInputs[2].value)){
        AddWrong("Email имеет неправильный формат" , 2);
    }
    if (allInputs[2].value.length > 30){
        AddWrong("" , 3);
        AddWrong("Email не должен превышать 64 символа" , 4);
    }

    // Возраст более 18 лет
    var date = allInputs[5].value.split("-");
    var year = parseInt(date[0]);
    var month = parseInt(date[1]);
    var day = parseInt(date[2]);
    var isOlder = false;
    if (year < (now.getFullYear() - 18)){
        isOlder = true;
    }
    if (year == (now.getFullYear() - 18)){
        if(month < (now.getMonth() + 1)){
            isOlder = true;
        }
    }
    if (year == (now.getFullYear() - 18)){
        if(month == (now.getMonth() + 1)){
            if(day < now.getDate()){
                isOlder = true;
            }
        }
    }
    if (isOlder == false){
        AddWrong("Вы должны быть старше 18 лет" , 5);
    }

    // Все поля заполненны
    for (let i = 0 ; i < allInputs.length; i++){
        if (allInputs[i].value === ''){
            AddWrong("Все поля должны быть заполнены!", i);
        }
    }

    if (wrongs > 0){ // Если ошибок ввода больше, чем 0, то false
        return false;
    } else{
        return true;
    }
}
            