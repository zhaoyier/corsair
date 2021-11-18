import { Modal } from "antd";
import * as React from "react";

interface Props {
	visible: boolean;
	urls: string[];
	onCancel: () => void;
}

interface State {
	activeUrl: string;
}

export class ImgPreviewModal extends React.Component<Props, State> {
	constructor(props) {
		super(props);
		this.state = {
			activeUrl: this.props.urls.length > 0 ? this.props.urls[0] : ""
		};
	}

	onChangePreviewImage = (active: string) => {
		this.setState({
			activeUrl: active
		});
	};

	componentDidUpdate(prevProps) {
		const { urls } = this.props;
		if (this.props.urls !== prevProps.urls) {
			this.setState({
				activeUrl: urls ? urls[0] : ""
			});
		}
	}

	render() {
		const { visible, urls, onCancel } = this.props;
		const { activeUrl } = this.state;
		return (
			<Modal key={"ImgModal"} visible={visible} footer={null} onCancel={onCancel} width={575}>
				<img
					alt=""
					style={{ width: 500, height: 500, objectFit: "contain" }}
					src={activeUrl}
				/>
				{urls.length > 1 &&
					urls.map((url, index) => (
						<a key={index} onClick={() => this.onChangePreviewImage(url)}>
							<img
								style={{
									width: 50,
									height: 50,
									objectFit: "cover",
									marginLeft: 10,
									marginTop: 20
								}}
								src={url}
							/>
						</a>
					))}
			</Modal>
		);
	}
}
