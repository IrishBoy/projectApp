const fetchSlogan = async () => {
    console.log("Initiating fetch request to backend.dddd..");
    try {
        const response = await fetch('http://localhost:8080/api/home');
        if (!response.ok) {
            console.error('Network response was not ok');
        }
        const slogan = await response.text();
        console.log("Slogan received:", slogan);
        const sloganElement = document.getElementById('slogan');
        if (sloganElement) {
            sloganElement.textContent = slogan;
        }
    } catch (error) {
        console.error('There has been a problem with your fetch operation:', error);
    }
};

fetchSlogan();