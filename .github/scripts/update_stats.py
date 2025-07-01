import os
import re

from pathlib import Path
from typing import Dict


ROOT = Path('.')
README = Path('README.md')
STATS_PATTERN = r'## Progress\s*\n\s*\| Difficulty \| Count \|\s*\n\s*\| ---------- \| ----- \|\s*\n\s*\| <span style="color:green">Easy</span>\s*\| \d+ \|\s*\n\s*\| <span style="color:orange">Medium</span>\s*\| \d+ \|\s*\n\s*\| <span style="color:red">Hard</span>\s*\| \d+ \|\s*\n\s*\| \*\*Total\*\*\s*\| \*\*\d+\*\* \|'


def count_problems() -> Dict[str, int]:
    ROOT = Path('.')

    counts = { 'easy': 0, 'medium': 0, 'hard': 0 }

    for difficulty in counts.keys():
        difficulty_dir = ROOT / difficulty

        if difficulty_dir.exists() and difficulty_dir.is_dir():
            counts[difficulty] = len(next(os.walk(difficulty_dir))[1])

    return counts


def update_readme(counts: Dict[str, int]) -> None:
    with open(README, 'r') as file:
        content = file.read()

    stats_table = f'''## Progress

| Difficulty | Count |
| ---------- | ----- |
| <span style="color:green">Easy</span>       | {counts["easy"]} |
| <span style="color:orange">Medium</span>     | {counts["medium"]} |
| <span style="color:red">Hard</span>       | {counts["hard"]} |
| **Total**  | **{sum(counts.values())}** |'''

    with open(README, 'r+') as file:
        if re.search(STATS_PATTERN, content):
            file.write(re.sub(STATS_PATTERN, stats_table, content))
        else:
            file.write(stats_table + '\n' + content)


if __name__ == '__main__':
    update_readme(count_problems())
