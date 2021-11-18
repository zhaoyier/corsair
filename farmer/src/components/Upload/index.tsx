import { message, Button, Upload as UploadAntd } from "antd";
import { UploadOutlined } from "@ant-design/icons";
import { UploadChangeParam } from "antd/lib/upload";
import * as React from "react";
import { getToken } from "utils/getToken";

export interface UploadResp {
	key: string;
	url: string;
	fileName?: string;
}

interface UploadProps {
	hideFileName?: boolean;
	text?: string;
	accept?: string;
	onChange: (url: UploadResp) => void;
}

interface UploadState {
	prefixURL: string;
	token: string;
	uploadUrl: string;
	uploading: boolean;
	fileName: string;
}

// const uploadURL = uploadQiniuUrl();

export class Upload extends React.PureComponent<UploadProps, UploadState> {
	static defaultProps: Partial<UploadProps> = {
		text: "Click to upload"
	};

	constructor(props) {
		super(props);

		this.state = {
			prefixURL: "",
			token: "",
			uploadUrl: "",
			uploading: false,
			fileName: ""
		};
	}

	onChange = (info: UploadChangeParam) => {
		const { onChange } = this.props;
		const { status, name, response } = info.file;
		const { prefixURL } = this.state;

		if (status === "done") {
			this.setState({
				uploading: false,
				fileName: name
			});
			const url = `${prefixURL}${response.key}`;
			onChange({
				url,
				key: response.key,
				fileName: name
			});
			message.success(`${name} upload success.`);
		} else if (status === "error") {
			this.setState({
				uploading: false
			});
			message.error(`${name} upload fail.`);
		} else if (status === "uploading") {
			this.setState({
				uploading: true
			});
		}
	};

	componentDidMount() {
		getToken(({ baseUrl, token, uploadUrl }) => {
			this.setState({
				prefixURL: baseUrl,
				token,
				uploadUrl
			});
		});
	}

	render() {
		const { text, hideFileName = false, accept = "" } = this.props;
		const { token, uploading, uploadUrl, fileName } = this.state;

		return (
			<UploadAntd
				action={uploadUrl}
				data={{ token }}
				showUploadList={false}
				accept={accept}
				onChange={this.onChange}>
				<Button loading={uploading}>
					<UploadOutlined />
					{text}
				</Button>
				{!hideFileName && fileName && (
					<span style={{ marginLeft: 10, fontSize: 12 }}>{fileName}</span>
				)}
			</UploadAntd>
		);
	}
}
