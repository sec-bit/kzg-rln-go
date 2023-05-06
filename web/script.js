async function connectToMetaMask() {
    try {
        // Request account access
        const accounts = await window.ethereum.request({ method: "eth_requestAccounts" });

        // Return the first account's address
        return accounts[0];
    } catch (error) {
        console.error("Error connecting to MetaMask:", error);
        alert("Error connecting to MetaMask. Please make sure it is installed and unlocked.");
    }
}

document.getElementById("connectWalletButton").addEventListener("click", async () => {
    const walletAddress = await connectToMetaMask();
    if (walletAddress) {
        document.getElementById("walletAddress").textContent = walletAddress;
        document.getElementById("connectWalletButton").disabled = true;
    }
});

async function sendTransactionWithMetaMask() {
    const pubKeyProofStr = document.getElementById("keyProof").value.trim();
    const result = JSON.parse(window.registerPubKeyOnChain(pubKeyProofStr));

    const target = result.target;
    const calldata = result.calldata;

    // Check if connected to MetaMask
    if (!window.ethereum || !window.ethereum.selectedAddress) {
        alert("Please connect to MetaMask first.");
        return;
    }

    // Get the current account
    const fromAddress = window.ethereum.selectedAddress;

    // Set up the transaction parameters
    const transactionParameters = {
        from: fromAddress,
        to: target,
        data: calldata,
        value: "0xde0b6b3a7640000",
    };

    // Send the transaction using MetaMask
    try {
        const txHash = await window.ethereum.request({
            method: "eth_sendTransaction",
            params: [transactionParameters],
        });
        console.log(`Transaction sent: ${txHash}`);
    } catch (error) {
        console.error("Error sending transaction:", error);
    }
}

function onSendButtonClick() {
    const messageInput = document.getElementById("messageInput");
    const message = messageInput.value.trim();
    if (message) {
        var msgWithProof = window.sendMessage(message);
        messageInput.value = "";
        window.receiveMessage(msgWithProof);
    }
}

document.getElementById("messageInput").addEventListener("keydown", (event) => {
    if (event.key === "Enter") {
        event.preventDefault();
        onSendButtonClick();
    }
});

document.getElementById("sendButton").addEventListener("click", onSendButtonClick);

function onGeneratePolyButtonClick() {
    const coefficientsInput = document.getElementById("coefficients");
    const coefficients = coefficientsInput.value.trim();
    if (!coefficients) {
        window.genNewPoly();
        // document.getElementById("registerOnChainButton").disabled = false;
        document.getElementById("registerOnServerButton").disabled = false;
        document.getElementById("generatePolyButton").disabled = true;
    }
}

document.getElementById("generatePolyButton").addEventListener("click", onGeneratePolyButtonClick);

document.getElementById("coefficients").addEventListener("input", () => {
    const coefficientsInput = document.getElementById("coefficients");
    const coefficients = coefficientsInput.value.trim();
    // document.getElementById("registerOnChainButton").disabled = !coefficients;
    document.getElementById("registerOnServerButton").disabled = !coefficients;
});

document.getElementById("registerOnChainButton").addEventListener("click", sendTransactionWithMetaMask);

function onRegisterOnServerButtonClick() {
    const commitmentInput = document.getElementById("commitment");
    const keyProofInput = document.getElementById("keyProof");
    window.registerOnServer(commitmentInput.value.trim(), keyProofInput.value.trim());
}
document.getElementById("registerOnServerButton").addEventListener("click", onRegisterOnServerButtonClick);

const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});
