import NotFoundSvg from "@/assets/illustrations/404.svg";
import { Button, Image, Typography } from "@douyinfe/semi-ui";
import { useNavigate, useRouter } from "@tanstack/react-router";
import { FC, PropsWithChildren } from "react";
import classes from "./pages.module.scss";

type INotFound = {};
const NotFound: FC<PropsWithChildren<INotFound>> = (props) => {
  console.log("notFound render...");
  const nav = useNavigate();
  const router = useRouter();
  const onBack = () => router.history.back();

  return (
    <section className={classes.notFoundBox}>
      <div className={classes.notFound}>
        <Image
          preview={false}
          src={NotFoundSvg}
          alt={`you are in not-found page`}
          width={300}
        />
        <Typography.Title heading={4}>你来到了一片荒芜</Typography.Title>
        <div className={classes.notFoundAction}>
          <Button onClick={() => nav({ to: "/" })} type={`tertiary`}>
            返回首页
          </Button>
          <Button onClick={onBack}>返回上页</Button>
        </div>
      </div>
    </section>
  );
};

export default NotFound;
