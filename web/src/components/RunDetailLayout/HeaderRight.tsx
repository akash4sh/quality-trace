import {Button} from 'antd';
import RunActionsMenu from 'components/RunActionsMenu';
import TestState from 'components/TestState';
import TraceActions from 'components/TraceActions';
import {TestState as TestStateEnum} from 'constants/TestRun.constants';
import {useTestSpecs} from 'providers/TestSpecs/TestSpecs.provider';
import {useTestRun} from 'providers/TestRun/TestRun.provider';
import {useTest} from 'providers/Test/Test.provider';
import * as S from './RunDetailLayout.styled';

interface IProps {
  testId: string;
  testVersion: number;
}

const HeaderRight = ({testId, testVersion}: IProps) => {
  const {isDraftMode} = useTestSpecs();
  const {run} = useTestRun();
  const {onRun} = useTest();
  const state = run.state;

  return (
    <S.Section $justifyContent="flex-end">
      {isDraftMode && <TraceActions />}
      {!isDraftMode && state && state !== TestStateEnum.FINISHED && (
        <S.StateContainer data-cy="test-run-result-status">
          <S.StateText>Test status:</S.StateText>
          <TestState testState={state} />
        </S.StateContainer>
      )}
      {!isDraftMode && state && state === TestStateEnum.FINISHED && (
        <Button data-cy="run-test-button" ghost onClick={onRun} type="primary">
          Run Test
        </Button>
      )}
      <RunActionsMenu isRunView resultId={run.id} testId={testId} testVersion={testVersion} />
    </S.Section>
  );
};

export default HeaderRight;