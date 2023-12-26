// Code generated by "stringer -type=actionType"; DO NOT EDIT.

package fzf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[actIgnore-0]
	_ = x[actStart-1]
	_ = x[actClick-2]
	_ = x[actInvalid-3]
	_ = x[actChar-4]
	_ = x[actMouse-5]
	_ = x[actBeginningOfLine-6]
	_ = x[actAbort-7]
	_ = x[actAccept-8]
	_ = x[actAcceptNonEmpty-9]
	_ = x[actAcceptOrPrintQuery-10]
	_ = x[actBackwardChar-11]
	_ = x[actBackwardDeleteChar-12]
	_ = x[actBackwardDeleteCharEof-13]
	_ = x[actBackwardWord-14]
	_ = x[actCancel-15]
	_ = x[actChangeBorderLabel-16]
	_ = x[actChangeHeader-17]
	_ = x[actChangePreviewLabel-18]
	_ = x[actChangePrompt-19]
	_ = x[actChangeQuery-20]
	_ = x[actClearScreen-21]
	_ = x[actClearQuery-22]
	_ = x[actClearSelection-23]
	_ = x[actClose-24]
	_ = x[actDeleteChar-25]
	_ = x[actDeleteCharEof-26]
	_ = x[actEndOfLine-27]
	_ = x[actForwardChar-28]
	_ = x[actForwardWord-29]
	_ = x[actKillLine-30]
	_ = x[actKillWord-31]
	_ = x[actUnixLineDiscard-32]
	_ = x[actUnixWordRubout-33]
	_ = x[actYank-34]
	_ = x[actBackwardKillWord-35]
	_ = x[actSelectAll-36]
	_ = x[actDeselectAll-37]
	_ = x[actToggle-38]
	_ = x[actToggleSearch-39]
	_ = x[actToggleAll-40]
	_ = x[actToggleDown-41]
	_ = x[actToggleUp-42]
	_ = x[actToggleIn-43]
	_ = x[actToggleOut-44]
	_ = x[actToggleTrack-45]
	_ = x[actToggleHeader-46]
	_ = x[actTrack-47]
	_ = x[actDown-48]
	_ = x[actUp-49]
	_ = x[actPageUp-50]
	_ = x[actPageDown-51]
	_ = x[actPosition-52]
	_ = x[actHalfPageUp-53]
	_ = x[actHalfPageDown-54]
	_ = x[actOffsetUp-55]
	_ = x[actOffsetDown-56]
	_ = x[actJump-57]
	_ = x[actJumpAccept-58]
	_ = x[actPrintQuery-59]
	_ = x[actRefreshPreview-60]
	_ = x[actReplaceQuery-61]
	_ = x[actToggleSort-62]
	_ = x[actShowPreview-63]
	_ = x[actHidePreview-64]
	_ = x[actTogglePreview-65]
	_ = x[actTogglePreviewWrap-66]
	_ = x[actTransform-67]
	_ = x[actTransformBorderLabel-68]
	_ = x[actTransformHeader-69]
	_ = x[actTransformPreviewLabel-70]
	_ = x[actTransformPrompt-71]
	_ = x[actTransformQuery-72]
	_ = x[actPreview-73]
	_ = x[actChangePreview-74]
	_ = x[actChangePreviewWindow-75]
	_ = x[actPreviewTop-76]
	_ = x[actPreviewBottom-77]
	_ = x[actPreviewUp-78]
	_ = x[actPreviewDown-79]
	_ = x[actPreviewPageUp-80]
	_ = x[actPreviewPageDown-81]
	_ = x[actPreviewHalfPageUp-82]
	_ = x[actPreviewHalfPageDown-83]
	_ = x[actPrevHistory-84]
	_ = x[actPrevSelected-85]
	_ = x[actPut-86]
	_ = x[actNextHistory-87]
	_ = x[actNextSelected-88]
	_ = x[actExecute-89]
	_ = x[actExecuteSilent-90]
	_ = x[actExecuteMulti-91]
	_ = x[actSigStop-92]
	_ = x[actFirst-93]
	_ = x[actLast-94]
	_ = x[actReload-95]
	_ = x[actReloadSync-96]
	_ = x[actDisableSearch-97]
	_ = x[actEnableSearch-98]
	_ = x[actSelect-99]
	_ = x[actDeselect-100]
	_ = x[actUnbind-101]
	_ = x[actRebind-102]
	_ = x[actBecome-103]
	_ = x[actResponse-104]
}

const _actionType_name = "actIgnoreactStartactClickactInvalidactCharactMouseactBeginningOfLineactAbortactAcceptactAcceptNonEmptyactAcceptOrPrintQueryactBackwardCharactBackwardDeleteCharactBackwardDeleteCharEofactBackwardWordactCancelactChangeBorderLabelactChangeHeaderactChangePreviewLabelactChangePromptactChangeQueryactClearScreenactClearQueryactClearSelectionactCloseactDeleteCharactDeleteCharEofactEndOfLineactForwardCharactForwardWordactKillLineactKillWordactUnixLineDiscardactUnixWordRuboutactYankactBackwardKillWordactSelectAllactDeselectAllactToggleactToggleSearchactToggleAllactToggleDownactToggleUpactToggleInactToggleOutactToggleTrackactToggleHeaderactTrackactDownactUpactPageUpactPageDownactPositionactHalfPageUpactHalfPageDownactOffsetUpactOffsetDownactJumpactJumpAcceptactPrintQueryactRefreshPreviewactReplaceQueryactToggleSortactShowPreviewactHidePreviewactTogglePreviewactTogglePreviewWrapactTransformactTransformBorderLabelactTransformHeaderactTransformPreviewLabelactTransformPromptactTransformQueryactPreviewactChangePreviewactChangePreviewWindowactPreviewTopactPreviewBottomactPreviewUpactPreviewDownactPreviewPageUpactPreviewPageDownactPreviewHalfPageUpactPreviewHalfPageDownactPrevHistoryactPrevSelectedactPutactNextHistoryactNextSelectedactExecuteactExecuteSilentactExecuteMultiactSigStopactFirstactLastactReloadactReloadSyncactDisableSearchactEnableSearchactSelectactDeselectactUnbindactRebindactBecomeactResponse"

var _actionType_index = [...]uint16{0, 9, 17, 25, 35, 42, 50, 68, 76, 85, 102, 123, 138, 159, 183, 198, 207, 227, 242, 263, 278, 292, 306, 319, 336, 344, 357, 373, 385, 399, 413, 424, 435, 453, 470, 477, 496, 508, 522, 531, 546, 558, 571, 582, 593, 605, 619, 634, 642, 649, 654, 663, 674, 685, 698, 713, 724, 737, 744, 757, 770, 787, 802, 815, 829, 843, 859, 879, 891, 914, 932, 956, 974, 991, 1001, 1017, 1039, 1052, 1068, 1080, 1094, 1110, 1128, 1148, 1170, 1184, 1199, 1205, 1219, 1234, 1244, 1260, 1275, 1285, 1293, 1300, 1309, 1322, 1338, 1353, 1362, 1373, 1382, 1391, 1400, 1411}

func (i actionType) String() string {
	if i < 0 || i >= actionType(len(_actionType_index)-1) {
		return "actionType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _actionType_name[_actionType_index[i]:_actionType_index[i+1]]
}
