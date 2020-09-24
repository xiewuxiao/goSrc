SELECT * FROM t_target t WHERE t.`secondtarget`= '' OR t.`firsttarget`='' ORDER BY targetappId
SELECT * FROM t_targetapp WHERE guid NOT IN  (
SELECT guid FROM t_targetapp WHERE totalamt>0 OR totalamt<0) ORDER BY id DESC

UPDATE t_targetapp SET totalamt=REPLACE(totalamt,'万元','');
UPDATE t_targetapp SET totalamt_cz=REPLACE(totalamt_cz,'万元','');
UPDATE t_targetapp SET totalamt=REPLACE(totalamt,'万','');
UPDATE t_targetapp SET totalamt_cz=REPLACE(totalamt_cz,'万','');
DELETE FROM t_target t WHERE t.`firsttarget` LIKE '%备注%'