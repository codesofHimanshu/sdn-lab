let socket;
let username = "";

const joinBtn = document.getElementById("joinBtn");
const sendBtn = document.getElementById("sendBtn");

const usernameInput =
    document.getElementById("usernameInput");

const messageInput =
    document.getElementById("messageInput");

const messagesDiv =
    document.getElementById("messages");

const chatArea =
    document.getElementById("chatArea");

joinBtn.addEventListener("click", () => {

    username = usernameInput.value.trim();

    if (!username) {
        alert("Enter username");
        return;
    }

    socket = new WebSocket(
        "wss://localhost:4433/ws"
    );

    socket.onopen = () => {

        socket.send(username);

        chatArea.classList.remove("hidden");

        joinBtn.disabled = true;
        usernameInput.disabled = true;

        addMessage(
            "Connected as " + username
        );
    };

    socket.onmessage = (event) => {

        addMessage(event.data);
    };

    socket.onerror = () => {

        alert("Connection failed");
    };
});

sendBtn.addEventListener("click", () => {

    const text =
        messageInput.value.trim();

    if (!text) return;

    socket.send(text);

    addMessage("YOU: " + text);

    messageInput.value = "";
});

function addMessage(text) {

    const div =
        document.createElement("div");

    div.className = "message";

    div.innerText = text;

    messagesDiv.appendChild(div);

    messagesDiv.scrollTop =
        messagesDiv.scrollHeight;
}