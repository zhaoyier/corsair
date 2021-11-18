import * as React from "react";
import * as moment from "moment";
import { setUserInfo, removeUserInfo } from "utils/user";
import { Form, Input, message, Button, Modal } from "antd";
import { redirectToLastPage } from "utils/auth";
import { UserOutlined, LockOutlined } from "@ant-design/icons";
import ChangePassword from "views/ChangePassword";
// import { Login } from "genServices/ezShipOMS/oms";
import { Login } from "genServices/fruit/orchard";
import { Base64 } from "utils/base64";

const styles = require("./login.scss");
const layout = {
	labelCol: { span: 0 },
	wrapperCol: { span: 24 }
};

const LoginForm: React.FC = () => {
	const [showResetPassword, setShowResetPassword] = React.useState(false);
	const [username, setUserName] = React.useState("");
	const [isLoading, setIsLoading] = React.useState(false);

	React.useEffect(() => {
		removeUserInfo();
	}, []);

	// 提交表单且数据验证成功后回调事件
	const handleSubmit = values => {
		setIsLoading(true);
		const base = new Base64();
		const { username, password } = values;
		// const url = "http://passport.65emall.net/user.User/login";
		Login({ username, password: base.encode(password), sourceCode: "web" }).then(
			resp => {
				setIsLoading(false);
				if (resp.code === "00000000") {
					message.success("login success!");
					const current = moment().unix();
					setUserInfo(
						Object.assign(
							{ ...resp.data, username },
							{ timeToExpire: current + resp.data.timeToExpire }
						)
					);
					redirectToLastPage();
					// } else if (resp.code === 6) {
					setUserName(username);
					// 	setShowResetPassword(true);
				} else {
					Modal.error({
						title: "Login error",
						content: resp.message
					});
				}
			},
			() => {
				setIsLoading(false);
				Modal.error({
					title: "rexxar Login error",
					content: "server error."
				});
			}
		);
	};

	const removePassowrdModal = () => {
		setShowResetPassword(false);
	};

	return (
		<React.Fragment>
			<div className={styles.FormBox}>
				<Form {...layout} onFinish={handleSubmit}>
					<Form.Item
						label="Username"
						name="username"
						rules={[{ required: true, message: "Please input your username!" }]}>
						<Input
							prefix={<UserOutlined className="site-form-item-icon" />}
							placeholder="Username"
						/>
					</Form.Item>
					<Form.Item
						label="Password"
						name="password"
						rules={[{ required: true, message: "Please input your password!" }]}>
						<Input
							prefix={<LockOutlined className="site-form-item-icon" />}
							type="password"
							placeholder="Password"
						/>
					</Form.Item>
					<Form.Item>
						<Button type="primary" loading={isLoading} htmlType="submit" block>
							登录
						</Button>
					</Form.Item>
				</Form>
				<Modal visible={showResetPassword} onCancel={removePassowrdModal} footer={null}>
					<p className={styles.errorTip}>!!密码已经过期，请重新设置密码后再使用</p>
					<p>密码必须同时包含：大写、小写英文字母、数字</p>
					<ChangePassword
						style={{ margin: "30px auto auto" }}
						payloadWhenPasswordExpire={{
							username,
							onResetPassword: () => removePassowrdModal()
						}}
					/>
				</Modal>
			</div>
		</React.Fragment>
	);
};

export default LoginForm;
