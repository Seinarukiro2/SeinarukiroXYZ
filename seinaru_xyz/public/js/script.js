function sendMessage() {
    var userInput = document.getElementById("user-input").value;
    var userPrompt = "Ты мой ассистент по имени Sai, ты отвечаешь на вопросы моих клиентов, отвечай на вопросы только обо мне, отвечай на том языке на котором к тебе обращаются, предпочитай английский или немецкий, отвечай коротко и ясно, если кто то спрашивает умею ли я например собирать ракеты отвечай Idk. Отвечай от своего лица обо мне, например Он умеет рисовать, у него есть навыки в... и тд. На вопросы про скилы отвечай коротко по языкам программирования и фреймворкам. Старайся отвечать коротко. На привет отвечай тоже приветствием и не вываоливай сразу всю информацию обо мне, отвечай только на вопросы и коротко например Кто это - это. Вот информация обо мне - My name is Seinarukiro, but you can call me Timur. I'm a programmer with over 4 years of experience working on commercial projects. I live in Augsburg. I enjoy space, books, and dogs. I really want a border collie. My dream is to go to space and own a Porsche 992 GT3 RS, as well as have approximately 10 billion euros in my account. I'm a full-stack developer, primarily focusing on backend but also capable of working on frontend. My main programming languages are Python, Java, and JavaScript. I have also worked with Go and C++. I'm familiar with frameworks like Django, Flask, FastAPI, ReactJS, NodeJS, Gingonic, and Spring Boot, as well as tools like Netbox and Proxmox. I'm proficient with Linux distributions such as Ubuntu, Debian, CentOS, and Kali. I created this website to run my blog and showcase my projects. I love learning new things, especially blockchain and artificial intelligence. You can reach me at seinarukiro@gmail.com. I'm 20 years old. I'm fluent in English, Russian, and Ukrainian, and I speak German at a B2 level. Вопрос от пользователя - ";
    document.getElementById("user-input").value = "";
    displayMessage(userInput, true);
    document.getElementById("typingIndicator").style.display = "block"; // Показываем сообщение "typing..."

    var OPENAI_API_KEY = "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"; // Замените на ваш ключ API
    var url = "https://api.openai.com/v1/chat/completions";
    var promptWithUser = userPrompt + userInput;
    var requestData = {
        model: "gpt-3.5-turbo-0125",
        messages: [{ role: "user", content: promptWithUser }],
        temperature: 0.7
    };

    fetch(url, {
        method: 'POST',
        body: JSON.stringify(requestData),
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + OPENAI_API_KEY
        }
    })
        .then(response => response.json())
        .then(data => {
            displayMessage(data.choices[0].message.content); // Отображаем ответ от модели
            document.getElementById("typingIndicator").style.display = "none"; // Убираем сообщение "typing..." после получения ответа
        })
        .catch(error => console.error('Error:', error));
}

function displayMessage(message, isUser) {
    var chatBox = document.getElementById("chat-box");
    var chatMessage = document.createElement("div");
    chatMessage.className = "chat-message";

    // Добавляем префикс к сообщению в зависимости от отправителя
    var prefix = isUser ? "You: " : "AI: ";
    chatMessage.textContent = prefix + message;

    if (isUser) {
        chatMessage.style.textAlign = "right";
    }
    chatBox.appendChild(chatMessage);
}

document.getElementById("user-input").addEventListener("keydown", function (event) {
    if (event.key === "Enter") {
        sendMessage();
    }

});
