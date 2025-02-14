
SELECT
	  org_id
	, id
	, uuid
	, is_instance_admin
	, first_name
	, last_name
	, email
	, username
	, password
	, created
	, created_by
	, modified
	, modified_by

FROM
	account.account

WHERE
	    org_id=$1
	AND id=ANY(string_to_array($2, ',')::bigint[])