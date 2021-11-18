import * as React from "react";
import { Upload, Button, message } from "antd";
import { UploadProps } from "antd/lib/upload";
import { getToken } from "utils/getToken";

interface UploadImgProps {
	onUploadSuccess: (url: string) => void;
	showIMG?: boolean; // 上传成功后显示图片
	baseURL?: string;
	maxSize?: number; // 限制上传文件的最大size
}

interface UploadState {
	token: string;
	baseUrl: string;
	loading: boolean;
	imgURL: string;
	uploadUrl: string;
}

export default class UploadImg extends React.PureComponent<
	UploadImgProps & UploadProps,
	UploadState
> {
	constructor(props) {
		super(props);
		this.state = {
			token: "",
			baseUrl: "",
			loading: false,
			imgURL: "",
			uploadUrl: ""
		};
	}

	fileChange(info) {
		if (info.file.response !== undefined) {
			const { baseURL, onUploadSuccess } = this.props;
			this.setState({
				loading: false
			});

			const base = baseURL ? baseURL : this.state.baseUrl;

			let imgurl = `${base}${info.file.response.key}`;
			onUploadSuccess(imgurl);
			this.setState({
				imgURL: imgurl
			});
		}
	}

	componentDidMount() {
		getToken(({ baseUrl, token, uploadUrl }) => {
			this.setState({
				baseUrl,
				token,
				uploadUrl
			});
		});
	}

	render() {
		const { loading, imgURL, uploadUrl } = this.state;
		const props = {
			showUploadList: false,
			action: uploadUrl,
			data: {
				token: this.state.token
			},
			onChange: this.fileChange.bind(this),
			beforeUpload: file => {
				const { maxSize } = this.props;
				if (maxSize && file.size > maxSize) {
					message.error(`文件过大：超过${maxSize / 1024 / 1024}M`);
					return false;
				}
				this.setState({ loading: true });
				return true;
			}
		};

		const { showIMG, children } = this.props;
		return (
			<span>
				<Upload {...props}>
					<Button type="primary" loading={loading}>
						{children || "上传图片"}
					</Button>
				</Upload>
				{showIMG && imgURL ? (
					<img src={imgURL} style={{ marginTop: 10, width: 120, display: "block" }} />
				) : null}
			</span>
		);
	}
}
export { uploadBase64 } from "./putb64";
