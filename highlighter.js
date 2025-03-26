// JSON syntax highlighter
document.addEventListener('DOMContentLoaded', function() {
    // Select all pre elements containing JSON
    const jsonElements = document.querySelectorAll('pre');
    
    jsonElements.forEach(function(element) {
        let content = element.innerHTML;
        
        // Highlight "Raw" and "RawV2" keys and their values in darker red with larger font
        content = content.replace(/"(Raw|RawV2)":\s*"([^"]*)"/g, 
            '<span style="color: #990000; font-weight: bold; font-size: 1.2em;">"$1": "$2"</span>');
        
        // Highlight "file" keys and values in blue with larger font
        content = content.replace(/"(file)":\s*"([^"]*)"/g, 
            '<span style="color: blue; font-weight: bold; font-size: 1.2em;">"$1": "$2"</span>');
        
        // Highlight "repository" keys and values in green with larger font
        content = content.replace(/"(repository)":\s*"([^"]*)"/g, 
            '<span style="color: #006600; font-weight: bold; font-size: 1.2em;">"$1": "$2"</span>');
        
        // Add emojis to "Verified" boolean values
        content = content.replace(/"(Verified)":\s*false/g, 
            '"$1": false ❌');
        
        content = content.replace(/"(Verified)":\s*true/g, 
            '"$1": true ✅');
        
        // Update the element with highlighted content
        element.innerHTML = content;
    });
});
