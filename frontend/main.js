document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('analyze-form');
    const resultsDiv = document.getElementById('results');

    form.addEventListener('submit', async (event) => {
        event.preventDefault();
        let url = document.getElementById('url').value;

        // Add protocol if missing
        if (!url.startsWith('http://') && !url.startsWith('https://')) {
            url = 'http://' + url;
        }

        const response = await fetch(`/v1/analyze?url=${encodeURIComponent(url)}`, {
            method: 'GET',
        });

        const data = await response.json();

        resultsDiv.innerHTML = `
            <h2>Analysis Results</h2>
            <pre>${JSON.stringify(data, null, 2)}</pre>
        `;
    });
});
