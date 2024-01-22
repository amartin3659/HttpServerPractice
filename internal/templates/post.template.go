package templates

import "fmt"

func PostTemplate(postID string) string {
  return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Update Post</title>
</head>
<body>
  <form method="POST" action="/user/post/%s" novalidate enctype="application/x-www-form-urlencoded">
      <label for="title">Title:</lable>
      <br>
      <input type="text" name="title" />
      <br>
      <br>
      <label for="body">Body:</label>
      <br>
      <textarea name="body" rows="10" cols="30"></textarea>
      <br>
      <br>
      <button type="submit">Update</button>
      <br>
      <br>
    </form>
</body>
</html>
    `, postID)
}
