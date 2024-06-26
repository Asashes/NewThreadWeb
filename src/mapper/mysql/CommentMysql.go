package mapper

import (
	"NewThread/src/pojo"
	"errors"
)

type CommentMysql struct{}

func NewCommentMysql() *CommentMysql {
	return &CommentMysql{}
}

func (c *CommentMysql) CommentTopThreeMysql(articleid int) ([]pojo.Comment, error) {
	var m []pojo.Comment

	err := Db.Raw("SELECT tc.id, tc.content, tc.likeCount, tc.creatTime, tc.rootCommentId,tc.userid,tc.toCommentId FROM ("+
		"SELECT  * FROM t_comment WHERE rootCommentId  IS NULL AND articleid = ? AND isDelete = 0 "+
		"UNION "+
		"SELECT * FROM (SELECT * FROM t_comment WHERE articleid = ? AND isDelete = 0) AS c "+
		"WHERE (SELECT count(*) FROM t_comment WHERE rootCommentId=c.rootCommentId AND id "+
		"<= c.id)<=3 AND rootCommentId  = ANY (SELECT id FROM t_comment WHERE rootCommentId IS NULL)  ORDER BY rootCommentId "+
		") AS tc", articleid, articleid).Scan(&m).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *CommentMysql) CommentAllMysql(commentid int) ([]pojo.Comment, error) {
	var m []pojo.Comment

	err := Db.Raw("SELECT id, content, likeCount, creatTime, rootCommentId,userid,toCommentId FROM t_comment WHERE rootCommentId = ?", commentid).Scan(&m).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *CommentMysql) CommentUploadMysql(content string, userid int, articleid int, rootcommentid int, tocommentid int) error {
	m := Db.Exec("INSERT INTO t_comment VALUES (NULL,?,0,?,?,0,?,?,NOW(),NOW());", content, userid, articleid, rootcommentid, tocommentid)

	rowsaffected := m.RowsAffected
	if rowsaffected == 0 {
		return errors.New("Insert---Comment---Mesg---Error")
	}
	return nil
}
