package postgresql

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	tableBypassRLSAttr         = "bypass_row_level_security"
	tableConnLimitAttr         = "connection_limit"
	tableCreateDBAttr          = "create_database"
	tableCreateRoleAttr        = "create_table"
	tableEncryptedPassAttr     = "encrypted_password"
	tableInheritAttr           = "inherit"
	tableLoginAttr             = "login"
	tableNameAttr              = "name"
	tablePasswordAttr          = "password"
	tableReplicationAttr       = "replication"
	tableSkipDropRoleAttr      = "skip_drop_table"
	tableSkipReassignOwnedAttr = "skip_reassign_owned"
	tableSuperuserAttr         = "superuser"
	tableValidUntilAttr        = "valid_until"
	tableRolesAttr             = "tables"

	// Deprecated options
	tableDepEncryptedAttr = "encrypted"
)

func resourcePostgreSQLTable() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostgreSQLTableCreate,
		Read:   resourcePostgreSQLTableRead,
		Update: resourcePostgreSQLTableUpdate,
		Delete: resourcePostgreSQLTableDelete,
		Exists: resourcePostgreSQLTableExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			tableNameAttr: {
				Type:     schema.TypeString,
				Required: true,
				Description: "The name of the table",
			},
		},
	}
}

func resourcePostgreSQLTableCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourcePostgreSQLTableRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourcePostgreSQLTableUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourcePostgreSQLTableDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourcePostgreSQLTableExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return false, nil
}
