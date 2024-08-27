import { postApiUserLogin } from "@/api/userApi.ts";
import { onError, onSuccess } from "@/core/callback.ts";
import { useAuth } from "@/provider/AuthProvider.tsx";
import { Button, Divider, Form } from "@douyinfe/semi-ui";
import { useMutation } from "@tanstack/react-query";
import { createLazyRoute, useNavigate } from "@tanstack/react-router";
import classes from "./pages.module.scss";

const { Input } = Form;

const LoginPage = () => {
  console.log("LoginPage render...");

  const { toLogin } = useAuth();
  const nav = useNavigate();

  const { mutate: doLogin, isPending } = useMutation({
    mutationFn: postApiUserLogin,
  });

  const handleLogin = (values: any) => {
    doLogin(values, {
      onSuccess: onSuccess("登录成功", (rtn) => {
        toLogin(
          {
            ...rtn.user,
            token: rtn.token,
          },
          () => {
            setTimeout(() => {
              nav({ to: "/admin" });
            }, 0);
          },
        );
      }),
      onError: onError(),
    });
  };

  return (
    <section className={classes.loginPage}>
      <div className={classes.loginBox}>
        <h1 className={classes.loginTitle}>系统登录</h1>
        <Divider margin={10} />
        <Form onSubmit={(values) => handleLogin(values)}>
          <Input
            label={`用户名`}
            placeholder={`请输入用户名`}
            field={`username`}
            rules={[{ required: true, message: "必填" }]}
          />
          <Input
            label={`密码`}
            placeholder={`请输入密码`}
            field={`password`}
            rules={[{ required: true, message: "必填" }]}
            mode={`password`}
          />
          <Button
            type={`primary`}
            htmlType={`submit`}
            block
            loading={isPending}
          >
            登录
          </Button>
        </Form>
      </div>
    </section>
  );
};

const LoginPageRoute = createLazyRoute("/login")({
  component: LoginPage,
});

export default LoginPageRoute;
