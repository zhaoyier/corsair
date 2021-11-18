import { message, Button, Form, Input, Modal } from "antd";
import { ChangePassword as ChangePasswordService } from "genServices/AdminLoginService";
import * as React from "react";
import { getUserInfo, removeUserInfo } from "utils/user";

const styleSheet = require("./index.scss");

const FormItem = Form.Item;

interface ChangePassword {
	style?: React.CSSProperties;
	payloadWhenPasswordExpire?: {
		onResetPassword: () => void;
		username: string;
	};
}

const ChangePassword = (props: ChangePassword) => {
	const { style = {}, payloadWhenPasswordExpire } = props;
	const styles = {
		...{ width: 400, margin: "60px auto" },
		...style
	};

	const handleSubmit = values => {
		const { payloadWhenPasswordExpire } = props;
		const userInfo = payloadWhenPasswordExpire
			? { username: payloadWhenPasswordExpire.username }
			: getUserInfo();
		ChangePasswordService(userInfo.username, values.password, values.newPassword).then(resp => {
			if (resp.code === 0) {
				message.success("修改成功");
				removeUserInfo();
				setTimeout(() => {
					window.location.replace("/login.html");
				}, 2000);
			} else {
				Modal.error({
					title: resp.msg
				});
			}
		});
	};

	const checkPassword = (rule, value, callback) => {
		console.log(rule);
		const form = this.props.form;
		if (value && value !== form.getFieldValue("newPassword")) {
			callback("两次输入密码不一致");
		} else {
			callback();
		}
	};

	const checkNewPassword = (_, value, callback) => {
		const { getFieldValue } = this.props.form;
		if (value === getFieldValue("password")) {
			return callback("新密码和旧密码不能相同");
		}
		const rule = /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,}$/;
		const hasCapLetter = /[A-Z]+/;
		const hasLetter = /[a-z]+/;
		const hasNumber = /[0-9]+/;
		const valid =
			rule.test(value) &&
			hasLetter.test(value) &&
			hasCapLetter.test(value) &&
			hasNumber.test(value);
		if (!valid) {
			return callback("字母与数字的组合,必须包含大写小写数字,长度需超过8个字符");
		}
		callback();
	};

	return (
		<Form onFinish={handleSubmit} style={styles}>
			{payloadWhenPasswordExpire && (
				<div className={styleSheet.usernameBar}>
					<span>用户名：</span>
					<span>{payloadWhenPasswordExpire.username}</span>
				</div>
			)}
			<FormItem
				label="旧密码"
				name="password"
				rules={[{ required: true, message: "请输入旧密码" }]}>
				<Input type="password" />
			</FormItem>
			<FormItem
				label="新密码"
				name="newPassword"
				rules={[
					{
						required: true,
						validator: checkNewPassword
					}
				]}>
				<Input type="password" />
			</FormItem>
			<FormItem
				label="再次输入新密码"
				name="confirmPassword"
				rules={[
					{ required: true, message: "请再次输入新密码" },
					{
						validator: checkPassword
					}
				]}>
				<Input type="password" />
			</FormItem>
			<FormItem style={{ textAlign: "center" }}>
				<Button type="primary" htmlType="submit">
					提交
				</Button>
			</FormItem>
		</Form>
	);
};

export default ChangePassword;
