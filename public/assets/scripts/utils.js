window.scrollTo({ top: 0, behavior: 'smooth' })

window.Clipboard = (function(window, document, navigator) {
    var textArea,
        copy;

    function isOS() {return navigator.userAgent.match(/ipad|iphone/i)}

    function createTextArea(text) {
        textArea = document.createElement('textArea');
        textArea.value = text;
        textArea.style.position = "fixed"
        textArea.style.top = "0vh"
        textArea.style.left = "0vw"

        document.body.appendChild(textArea);
    }

    function selectText() {
        var range,
            selection;

        if (isOS()) {
            console.log("IS IOS");
            
            range = document.createRange();
            range.selectNodeContents(textArea);
            selection = window.getSelection();
            selection.removeAllRanges();
            selection.addRange(range);
            textArea.setSelectionRange(0, 999999);
        } else {
            textArea.select();
        }
    }

    function copyToClipboard() {
        let exec = document.execCommand('copy');
        console.log("copy res", exec);   
        document.body.removeChild(textArea);
        return exec
    }

    copy = async function(textTocopy) {
        try {
            createTextArea(textTocopy);
            selectText();
            copyToClipboard();
        } catch (error) {console.log("copy func err", error)}
    };
    
    return {copy: copy};
})(window, document, navigator)

const CopyToClipboard = Clipboard['copy']

window.Toast = (text, link = "#") => {
    Toastify({
        text: text,
        duration: 2000,
        destination: link,
        newWindow: true,
        close: true,
        gravity: "bottom", // `top` or `bottom`
        position: "center", // `left`, `center` or `right`
        stopOnFocus: true, // Prevents dismissing of toast on hover
        style: {
            background: "#21569e",
        },
        onClick: function () { } // Callback after click
    }).showToast();
}

const Toast = window.Toast