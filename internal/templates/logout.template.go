package templates

func GetLogout() string {
  return `
<!DOCTYPE html>
<html>
<head>
    <title>Logout</title>
</head>
<body>
  <form method="POST" action="/user/logout" novalidate>
      <button type="submit">Logout</button>
    </form>
</body>
</html>
  `
}
