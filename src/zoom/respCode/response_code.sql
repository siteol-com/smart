SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- Reset Table
TRUNCATE TABLE response_code;
-- Record Start
INSERT INTO response_code VALUES ('1', 'E000', '0', 'E', '系统异常', 'System exception', '系统异常（默认）', '1', '0', null, null);
INSERT INTO response_code VALUES ('2', 'E001', '0', 'E', '参数非法', 'Illegal parameters', '参数非法（默认）（免翻译）', '1', '0', null, null);
INSERT INTO response_code VALUES ('3', 'E002', '0', 'E', '尚未登陆', 'Not logged in yet', '尚未登陆（默认）', '1', '0', null, null);
INSERT INTO response_code VALUES ('4', 'E003', '0', 'E', '无权访问', 'No access rights', '无权访问（默认）', '1', '0', null, null);
INSERT INTO response_code VALUES ('5', 'E004', '0', 'E', '路径不存在', 'Path does not exist', '路径不存在（默认）', '1', '0', null, null);
INSERT INTO response_code VALUES ('6', 'S000', '0', 'S', '处理成功', 'Processed successfully', '处理成功（默认）', '1', '0', null, null);
INSERT INTO response_code VALUES ('7', 'S001', '0', 'S', '登陆成功', 'Landed successfully', '登陆成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('8', 'S002', '0', 'S', '密码重置成功', 'Password reset successful', '密码重置成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('9', 'F000', '0', 'F', '处理失败', 'Processing failed', '处理失败（默认）', '1', '0', null, null);
INSERT INTO response_code VALUES ('10', 'F001', '0', 'F', '登陆失败，请联系管理员', 'Login failed, please contact the administrator', '登陆失败，请联系管理员', '1', '0', null, null);
INSERT INTO response_code VALUES ('11', 'F002', '0', 'F', '异常登陆，请联系管理员', 'Abnormal login, please contact the administrator', '异常登陆，请联系管理员', '1', '0', null, null);
INSERT INTO response_code VALUES ('12', 'F003', '0', 'F', '密码重置失败，请联系管理员', 'Password reset failed, please contact the administrator', '密码重置失败，请联系管理员', '1', '0', null, null);
INSERT INTO response_code VALUES ('13', 'S100', '1', 'S', '字典创建成功', 'Dictionary creation successful', '字典创建成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('14', 'S101', '1', 'S', '字典编辑成功', 'Dictionary editing successful', '字典编辑成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('15', 'S102', '1', 'S', '字典排序成功', 'Dictionary sorting successful', '字典排序成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('16', 'S103', '1', 'S', '字典封存成功', 'Dictionary sealing successful', '字典封存成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('17', 'F100', '1', 'F', '字典查询失败', 'Dictionary query failed', '字典查询失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('18', 'F101', '1', 'F', '字典分组下字典值唯一', 'Dictionary value is unique under dictionary group', '字典分组下字典值唯一', '1', '0', null, null);
INSERT INTO response_code VALUES ('19', 'F102', '1', 'F', '内置字典禁止刪除', 'Built-in dictionary cannot be deleted', '内置字典禁止刪除', '1', '0', null, null);
INSERT INTO response_code VALUES ('20', 'F103', '1', 'F', '字典排序失败', 'Dictionary sorting failed', '字典排序失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('21', 'S200', '2', 'S', '响应码创建成功', 'Response code creation successful', '响应码创建成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('22', 'S201', '2', 'S', '响应码创建成功,实际响应码为{{code}}', 'Response code creation successful, actual response code is {{code}}', '响应码创建成功,实际响应码为{{code}}', '1', '0', null, null);
INSERT INTO response_code VALUES ('23', 'S202', '2', 'S', '响应码编辑成功', 'Response code editing successful', '响应码编辑成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('24', 'S203', '2', 'S', '响应码封存成功', 'Response code sealing successful', '响应码封存成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('25', 'F200', '2', 'F', '响应码查询失败', 'Response code query failed', '响应码查询失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('26', 'F201', '2', 'F', '响应码全局唯一', 'Response code is globally unique', '响应码全局唯一', '1', '0', null, null);
INSERT INTO response_code VALUES ('27', 'F202', '2', 'F', '内置响应码禁止删除', 'Built-in response code cannot be deleted', '内置响应码禁止删除', '1', '0', null, null);
INSERT INTO response_code VALUES ('28', 'S300', '3', 'S', '路由创建成功', 'Route creation successful', '路由创建成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('29', 'S301', '3', 'S', '路由编辑成功', 'Route editing successful', '路由编辑成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('30', 'S302', '3', 'S', '路由删除成功', 'Route deletion successful', '路由删除成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('31', 'F300', '3', 'F', '路由查询失败', 'Route query failed', '路由查询失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('32', 'F301', '3', 'F', '路由地址全局唯一', 'Route address is globally unique', '路由地址全局唯一', '1', '0', null, null);
INSERT INTO response_code VALUES ('33', 'F302', '3', 'F', '路由名称全局唯一', 'Route name is globally unique', '路由名称全局唯一', '1', '0', null, null);
INSERT INTO response_code VALUES ('34', 'F303', '3', 'F', '内置路由禁止删除', 'Built-in route cannot be deleted', '内置路由禁止删除', '1', '0', null, null);
INSERT INTO response_code VALUES ('35', 'F304', '3', 'F', '路由删除', 'Route deletion', '路由删除', '1', '0', null, null);
INSERT INTO response_code VALUES ('36', 'S400', '4', 'S', '权限创建成功', 'Permission creation successful', '权限创建成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('37', 'S401', '4', 'S', '权限编辑成功', 'Permission editing successful', '权限编辑成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('38', 'S402', '4', 'S', '权限删除成功', 'Permission deletion successful', '权限删除成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('39', 'S403', '4', 'S', '权限排序成功', 'Permissions sorted successfully', '权限排序成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('40', 'F400', '4', 'F', '权限查询失败', 'Permission query failed', '权限查询失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('41', 'F401', '4', 'F', '权限别名全局唯一', 'Permission alias is globally unique', '权限别名全局唯一', '1', '0', null, null);
INSERT INTO response_code VALUES ('42', 'F402', '4', 'F', '权限名称全局唯一', 'Permission name is globally unique', '权限名称全局唯一', '1', '0', null, null);
INSERT INTO response_code VALUES ('43', 'F403', '4', 'F', '内置权限禁止删除', 'Built-in permission cannot be deleted', '内置权限禁止删除', '1', '0', null, null);
INSERT INTO response_code VALUES ('44', 'F404', '4', 'F', '权限配置路由同步失败', 'Permission configuration route synchronization failed', '权限配置路由同步失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('45', 'F405', '4', 'F', '权限配置删除失败', 'Permission configuration deletion failed', '权限配置删除失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('46', 'S500', '5', 'S', '角色创建成功', 'Role creation successful', '角色创建成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('47', 'S501', '5', 'S', '角色编辑成功', 'Role editing successful', '角色编辑成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('48', 'S502', '5', 'S', '角色删除失败', 'Role deletion failed', '角色删除失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('49', 'F500', '5', 'F', '角色查询失败', 'Role query failed', '角色查询失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('50', 'F501', '5', 'F', '内置角色禁止编辑', 'Built-in roles cannot be edited', '内置角色禁止编辑', '1', '0', null, null);
INSERT INTO response_code VALUES ('51', 'F502', '5', 'F', '角色名全局唯一', 'Role name is globally unique', '角色名全局唯一', '1', '0', null, null);
INSERT INTO response_code VALUES ('52', 'F503', '5', 'F', '角色权限配置失败', 'Role permission configuration failed', '角色权限配置失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('53', 'F504', '5', 'F', '角色删除失败', 'Role deletion failed', '角色删除失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('54', 'S600', '6', 'S', '部门创建成功', 'Department created successfully', '集团部门创建成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('55', 'S601', '6', 'S', '部门编辑成功', 'Department editing successful', '集团部门编辑成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('56', 'S602', '6', 'S', '部门删除成功', 'Department deleted successfully', '集团部门删除成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('57', 'S603', '6', 'S', '部门排序成功', 'Department sorting successful', '集团部门排序成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('58', 'S604', '6', 'S', '部门迁移成功', 'Department migration successful', '集团部门迁移成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('59', 'F600', '6', 'F', '部门查询失败', 'Department query failed', '集团部门查询失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('60', 'F601', '6', 'F', '部门名称全局唯一', 'Department name Globally unique', '集团部门名称全局唯一', '1', '0', null, null);
INSERT INTO response_code VALUES ('61', 'F602', '6', 'F', '内置部门禁止删除', 'Deletion of built-in departments is prohibited', '内置集团部门禁止删除', '1', '0', null, null);
INSERT INTO response_code VALUES ('62', 'F603', '6', 'F', '部门删除失败', 'Department deletion failed', '集团部门删除失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('63', 'F604', '6', 'F', '部门存在子部门禁止删除', 'Department has sub-departments and is prohibited from deletion', '集团部门存在子部门禁止删除', '1', '0', null, null);
INSERT INTO response_code VALUES ('64', 'F605', '6', 'F', '部门存在成员禁止删除', 'Department members cannot be deleted if they exist', '集团部门存在成员禁止删除', '1', '0', null, null);
INSERT INTO response_code VALUES ('65', 'F606', '6', 'F', '部门迁移失败', 'Department migration failed', '集团部门迁移失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('66', 'S700', '7', 'S', '登陆账号创建成功', 'Login account created successfully', '登陆账号创建成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('67', 'S701', '7', 'S', '登陆账号编辑成功', 'Login account edited successfully', '登陆账号编辑成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('68', 'S702', '7', 'S', '登陆账号删除成功', 'Login account deleted successfully', '登陆账号删除成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('69', 'S703', '7', 'S', '登陆账号重置成功', 'Login account reset successfully', '登陆账号重置成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('70', 'F700', '7', 'F', '登陆账号查询失败', 'Login account query failed', '登陆账号查询失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('71', 'F701', '7', 'F', '登陆账号全局唯一', 'Login account is globally unique', '登陆账号全局唯一', '1', '0', null, null);
INSERT INTO response_code VALUES ('72', 'F702', '7', 'F', '账号角色同步失败', 'Account role synchronization failed', '账号角色同步失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('73', 'F703', '7', 'F', '特殊账号禁止编辑', 'Editing of special accounts is prohibited', '特殊账号禁止编辑', '1', '0', null, null);
INSERT INTO response_code VALUES ('74', 'F704', '7', 'F', '登陆账号删除失败', 'Login account deletion failed', '登陆账号删除失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('75', 'F705', '7', 'F', '登陆账号重置失败', 'Login account reset failed', '登陆账号重置失败', '1', '0', null, null);
INSERT INTO response_code VALUES ('76', 'S800', '8', 'S', '系统配置编辑成功', 'System configuration editing successful', '系统配置编辑成功', '1', '0', null, null);
INSERT INTO response_code VALUES ('77', 'F800', '8', 'F', '系统配置查询失败', 'System configuration query failed', '系统配置查询失败', '1', '0', null, null);
-- Record End
SET FOREIGN_KEY_CHECKS = 1;