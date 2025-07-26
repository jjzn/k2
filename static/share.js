async function share(url) {
    try {
        await navigator.share({ url });
    } catch (error) {
        try {
            await navigator.clipboard.writeText(url);
            alert('URL copied to clipboard');
        } catch (error) {
            alert(url);
        }
    }
}
