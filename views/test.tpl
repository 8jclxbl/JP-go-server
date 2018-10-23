<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="content-Type" content=utf-8>
        <title>BeegoLearning</title>
    </head>
    <body>
        <header>
            <div class="register">
                <form  method="post">
                    <input name="UserNickname" type="text" value="nick">用户昵称：</intput>
                    <br/>
                    <input name="UserName"type="text" value="a123">用户名：</intput>
                    <br/>
                    <input name="UserPass" type="text" value="123456">密码：</intput>
                    <br/>
                    <input name="UserSex"type="text" value="male">性别：</intput>
                    <br/>
                    <input name="UserBirthday" type="text" value="1990-01-01">生日：</intput>
                    <br/>
                    <input name="UserPhone"type="text" value="12345678901">手机号：</intput>
                    <br/>
                    <input name="UserEmail" type="text" value="123@123.com">电子邮箱：</intput>
                    <br/>
                    <input name="UserHomeplace"type="text" value="china">出生地：</intput>
                    <br/>
                    <input name="UserAddress" type="text" value="UserAddress">地址：</intput>
                    <br/>
                    <input name="UserImgurl"type="text" value="UserImgurl">头像：</intput>
                    <br/>
                    <input type="submit" value="提交">
                </form>
            </div>
        </header>
        <footer>
            <h1>{{.user}}</h1>
            <h1>{{.pw}}</h1>
        </footer>
    </body>
</html>