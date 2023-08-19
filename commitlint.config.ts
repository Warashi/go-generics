import type { UserConfig } from '@commitlint/types';
import { RuleConfigSeverity } from "@commitlint/types";

const config: UserConfig = {
  rules: {
    "body-max-line-length": [RuleConfigSeverity.Disabled],
    "header-max-length": [RuleConfigSeverity.Disabled],
    "subject-case": [RuleConfigSeverity.Disabled],
  },
  extends: ['@commitlint/config-conventional'],
}

export default config;
