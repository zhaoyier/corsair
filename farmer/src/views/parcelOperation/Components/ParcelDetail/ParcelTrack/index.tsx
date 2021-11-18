import * as React from "react";
import { observer } from "mobx-react-lite";
import { Timeline } from "antd";
import { useStores } from "views/parcelOperation/hooks";
import { formatUnixTime } from "utils/time";

const styles = require("./index.scss");

const ParcelTrack = () => {
	const { parcelDetailStore } = useStores();
	const { parcelTrack } = parcelDetailStore;
	return (
		<div className={styles.timeline}>
			<Timeline reverse>
				{parcelTrack &&
					parcelTrack.length > 0 &&
					parcelTrack.map((item, idx) => (
						<Timeline.Item key={idx}>
							{item.context}&emsp;&emsp;{formatUnixTime(item.createDate)}
						</Timeline.Item>
					))}
			</Timeline>
		</div>
	);
};

export default observer(ParcelTrack);
