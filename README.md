# **GoQuery Crawler**

A simple guide to building a web crawler using Go and GoQuery.

---

## **Setup Instructions**

1. **Initialize a New Module**  
   Run the following command to initialize a Go module:  
   ```bash
   go mod init <your-module-name>
   ```

2. **Download Dependencies**  
   Use `go mod tidy` to fetch and install the necessary dependencies:  
   ```bash
   go mod tidy
   ```

3. **Run the Application**  
   Execute the main Go file:  
   ```bash
   go run main.go
   ```

---

## **Using GoQuery for Web Scraping**

### **1. Inspect the HTML Structure**  
   - Before writing your scraper, inspect the HTML of the target website using your browser’s developer tools (`Ctrl+Shift+I` or right-click → "Inspect").
   - Identify the relevant **CSS classes**, **IDs**, or other HTML attributes you want to scrape.

### **2. Update the Code**  
   - After creating the GoQuery document (using `goquery.NewDocumentFromReader`), update the `.Find` selectors with the corresponding CSS classes, IDs, or element tags identified during inspection.  
   - Example:  
     ```go
     doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
         fmt.Println(s.Text())
     })
     ```

---

## **Example Workflow**

1. **Inspect the Target Website**:  
   Look for elements you want to scrape (e.g., titles, links, images).  
   
   Example:  
   ```html
   <div class="post-title">Scraped Title</div>
   ```

2. **Update the Selector in Code**:  
   Replace `.Find` values to match the desired elements.  
   Example:  
   ```go
   doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
       fmt.Println(s.Text())
   })
   ```

3. **Run the Scraper**:  
   Execute `main.go` and check the output.

---

### **Best Practices**
- Regularly inspect the target website for changes in HTML structure, as these may break your scraper.
- Use error handling in your code to manage issues like missing elements or HTTP request failures.

---

This guide should help you get started with creating a web crawler using GoQuery. Happy scraping!

