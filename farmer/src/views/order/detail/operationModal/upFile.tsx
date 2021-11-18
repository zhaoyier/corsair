import * as React from "react";
import { Row, Col } from "antd";
import UploadImg from "components/UploadImg/index";
interface UpFileProps {
	setUpFileLinkArr: (upFileLinkArr: string[]) => void;
}
export function UpFile(props: UpFileProps) {
	const [imgUrlList, setImgUrlList] = React.useState<string[]>([]);
	const { setUpFileLinkArr } = props;
	function successUpload(url) {
		const result = imgUrlList.concat(url);
		setImgUrlList(result);
		setUpFileLinkArr(result);
	}
	return (
		<Row>
			{imgUrlList.map((imgUrl, index) => (
				<Col span={8} key={imgUrl + index}>
					<img src={imgUrl} width={100} height={100} />
				</Col>
			))}
			{imgUrlList.length < 3 && (
				<UploadImg onUploadSuccess={successUpload} maxSize={5242880} />
			)}
		</Row>
	);
}
