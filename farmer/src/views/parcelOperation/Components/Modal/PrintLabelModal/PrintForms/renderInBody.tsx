import * as React from "react";
import * as ReactDOM from "react-dom";

interface RenderInBodyProps {
	className?: any;
	style?: React.CSSProperties;
	onClick?: () => void;
}

export default class RenderInBody extends React.Component<RenderInBodyProps, {}> {
	private popup: HTMLDivElement;

	componentDidMount() {
		this.popup = document.createElement("div");
		document.body.appendChild(this.popup);
		this.renderLayer();
	}

	componentDidUpdate() {
		this.renderLayer();
	}

	componentWillUnmount() {
		ReactDOM.unmountComponentAtNode(this.popup);
	}

	renderLayer = () => {
		ReactDOM.render(<div {...this.props}>{this.props.children}</div>, this.popup);
	};

	render() {
		// Render a placeholder
		return <div {...this.props} children={null} />;
	}
}
