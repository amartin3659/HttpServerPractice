package templates

func LoginTemplate() string {
	return `
<!DOCTYPE html>
<html>
<head>
    <title>Login</title>
</head>
<body>
  <form method="POST" action="/user/login" novalidate>
      <label for="email">Email:</lable>
      <br>
      <input type="email" name="email" />
      <br>
      <br>
      <label for="password">Password:</label>
      <br>
      <input type="password" name="password" />
      <br>
      <br>
      <button type="submit">Login</button>
      <br>
      <br>
    </form>
</body>
</html>
  `
}
