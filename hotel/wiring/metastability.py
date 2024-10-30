import matplotlib.pyplot as plt
import pandas as pd
from matplotlib.patches import Rectangle
import numpy as np
import os
from matplotlib.ticker import AutoMinorLocator, LogLocator, NullFormatter

def q99(x):
    return x.quantile(0.99)

markers = ['o', 's', '+', '^', '.', '*']
LEGEND_SIZE=15
LABEL_SIZE=18
TICK_SIZE=15
FIG_W_SIZE=6
FIG_H_SIZE=3
META_WIDTH = 5
y_minor = LogLocator(base = 10.0, subs = np.arange(1.0, 10.0) * 0.1, numticks = 10)
META_YLIM=(1,2000000)
META_WIDTH = 5

ok_df = pd.read_csv("stats_1.csv")
trigger_df = pd.read_csv("stats_2.csv")
not_ok_df = pd.read_csv("stats_3.csv")

min_start = np.min(ok_df.Start)
ok_df['start_rel'] = (ok_df.Start - min_start) / 1e9 # Convert to seconds
trigger_df['start_rel'] = (trigger_df.Start - min_start) / 1e9 # Convert to seconds
not_ok_df['start_rel'] = (not_ok_df.Start - min_start) / 1e9 # Convert to seconds
ok_df.Duration = ok_df.Duration / 1e6 # Convert to ms
not_ok_df.Duration = not_ok_df.Duration / 1e6 # Convert to ms
trigger_df.Duration = trigger_df.Duration / 1e6 # Convert to ms
ok_df['start_rel'] = ok_df['start_rel'].astype(int)
not_ok_df['start_rel'] = not_ok_df['start_rel'].astype(int)
trigger_df['start_rel'] = trigger_df['start_rel'].astype(int)
big_df = pd.concat([ok_df, trigger_df, not_ok_df])
filtered_df = big_df.groupby(big_df['start_rel']).agg(start_rel=('start_rel', 'mean'), avg=('Duration', 'mean'), p99=('Duration', q99))

trigger_start = np.min(trigger_df['start_rel'])
trigger_end = np.max(trigger_df['start_rel'])

fig, ax = plt.subplots(figsize=(META_WIDTH,FIG_H_SIZE))
ax.plot(filtered_df.start_rel, filtered_df.p99, label='99th %', color='black')
ax.plot(filtered_df.start_rel, filtered_df.avg, label='Average', color='gray')
ax.set_ylim(*META_YLIM)
ax.set_yscale('log')
ax.set_xlabel('Time (s)',fontsize=LABEL_SIZE)
ax.set_ylabel('Latency log(ms)', fontsize=LABEL_SIZE)
ax.tick_params(labelsize=TICK_SIZE)
ax.yaxis.set_minor_locator(y_minor)
ax.yaxis.set_minor_formatter(NullFormatter())
ax.grid(visible=True, which='major', axis='both', alpha=0.5)
ax.axvspan(trigger_start, trigger_end, color="red", alpha=0.1)
ax.text(30, 200, "Trigger", fontsize=LEGEND_SIZE)
plt.margins(x=0,y=0)
ax.legend(fontsize=LEGEND_SIZE)
plt.title("Type 1 Metastability Failure", fontsize=LABEL_SIZE)
fig.savefig('metastability.pdf', bbox_inches='tight',dpi=1200)
plt.show()