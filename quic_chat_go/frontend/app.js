let socket;
let username = "";

let localStream = null;

const loginBtn =
    document.getElementById("loginBtn");

const sendBtn =
    document.getElementById("sendBtn");

const startCameraBtn =
    document.getElementById("startCameraBtn");

const muteBtn =
    document.getElementById("muteBtn");

const cameraBtn =
    document.getElementById("cameraBtn");

const usernameInput =
    document.getElementById("usernameInput");

const passwordInput =
    document.getElementById("passwordInput");

const messageInput =
    document.getElementById("messageInput");

const messagesDiv =
    document.getElementById("messages");

const usersList =
    document.getElementById("usersList");

const loginArea =
    document.getElementById("loginArea");

const mainArea =
    document.getElementById("mainArea");

const localVideo =
    document.getElementById("localVideo");

const remoteVideo =
    document.getElementById("remoteVideo");

loginBtn.addEventListener("click", () => {

    username =
        usernameInput.value.trim();

    const password =
        passwordInput.value.trim();

    if (!username || !password) {

        alert("Enter credentials");

        return;
    }

    socket = new WebSocket(
        "wss://localhost:4433/ws"
    );

    socket.onopen = () => {

        socket.send(
            JSON.stringify({
                type: "login",
                username,
                password
            })
        );
    };

    socket.onmessage = (event) => {

        const data =
            JSON.parse(event.data);

        handleMessage(data);
    };
});

function handleMessage(data) {

    if (data.type === "login_success") {

        loginArea.classList.add("hidden");

        mainArea.classList.remove("hidden");

        addMessage(
            "Logged in as " + username
        );
    }

    if (data.type === "chat") {

        addMessage(data.message);
    }

    if (data.type === "users") {

        updateUsers(data.users);
    }
}

sendBtn.addEventListener("click", () => {

    const text =
        messageInput.value.trim();

    if (!text) return;

    socket.send(
        JSON.stringify({
            type: "chat",
            message: text
        })
    );

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

function updateUsers(users) {

    usersList.innerHTML = "";

    users.forEach(user => {

        const div =
            document.createElement("div");

        div.className = "user";

        div.innerText = user;

        usersList.appendChild(div);
    });
}

startCameraBtn.addEventListener(
    "click",
    async () => {

        try {

            localStream =
                await navigator
                .mediaDevices
                .getUserMedia({
                    video: true,
                    audio: true
                });

            localVideo.srcObject =
                localStream;

        } catch (err) {

            console.error(err);

            alert(
                "Camera access denied"
            );
        }
    }
);

muteBtn.addEventListener("click", () => {

    if (!localStream) return;

    localStream
        .getAudioTracks()
        .forEach(track => {

            track.enabled =
                !track.enabled;
        });
});

cameraBtn.addEventListener("click", () => {

    if (!localStream) return;

    localStream
        .getVideoTracks()
        .forEach(track => {

            track.enabled =
                !track.enabled;
        });
});