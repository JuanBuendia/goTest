/*==============================================================*/
/* Table: ASSETS                                                */
/*==============================================================*/
create table ASSETS (
   ID                   SERIAL               not null,
   KIN_ID               INT4                 null,
   NAME                 VARCHAR(200)         null,
   DESCRIPTION          VARCHAR(2000)        null,
   FILE_ID              VARCHAR(400)         null,
   constraint PK_ASSETS primary key (ID)
);

/*==============================================================*/
/* Index: ASSETS_PK                                             */
/*==============================================================*/
create unique index ASSETS_PK on ASSETS (
ID
);

/*==============================================================*/
/* Index: RELATIONSHIP_10_FK                                    */
/*==============================================================*/
create  index RELATIONSHIP_10_FK on ASSETS (
KIN_ID
);

/*==============================================================*/
/* Table: ASSTES_SITE                                           */
/*==============================================================*/
create table ASSTES_SITE (
   ID                   SERIAL               not null,
   ASS_ID               INT4                 null,
   SIT_ID               INT4                 null,
   constraint PK_ASSTES_SITE primary key (ID)
);

/*==============================================================*/
/* Index: RELATIONSHIP_6_FK                                     */
/*==============================================================*/
create  index RELATIONSHIP_6_FK on ASSTES_SITE (
ASS_ID
);

/*==============================================================*/
/* Index: RELATIONSHIP_7_FK                                     */
/*==============================================================*/
create  index RELATIONSHIP_7_FK on ASSTES_SITE (
SIT_ID
);

/*==============================================================*/
/* Table: CATEGORY                                              */
/*==============================================================*/
create table CATEGORY (
   ID                   SERIAL               not null,
   NAME                 VARCHAR(200)         null,
   DESCRIPTION          VARCHAR(1000)        null,
   THUMBNAIL            VARCHAR(300)         null,
   constraint PK_CATEGORY primary key (ID)
);

/*==============================================================*/
/* Index: CATEGORY_PK                                           */
/*==============================================================*/
create unique index CATEGORY_PK on CATEGORY (
ID
);

/*==============================================================*/
/* Table: COMMENTS                                              */
/*==============================================================*/
create table COMMENTS (
   ID                   SERIAL               not null,
   VIS_ID               INT4                 null,
   COMMENT              VARCHAR(500)         null,
   DATE                 DATE                 null,
   constraint PK_COMMENTS primary key (ID)
);

/*==============================================================*/
/* Index: COMMENTS_PK                                           */
/*==============================================================*/
create unique index COMMENTS_PK on COMMENTS (
ID
);

/*==============================================================*/
/* Table: KIND                                                  */
/*==============================================================*/
create table KIND (
   ID                   SERIAL               not null,
   NAME                 VARCHAR(200)         null,
   constraint PK_KIND primary key (ID)
);

/*==============================================================*/
/* Index: TYPE_PK                                               */
/*==============================================================*/
create unique index TYPE_PK on KIND (
ID
);

/*==============================================================*/
/* Table: PLACE                                                 */
/*==============================================================*/
create table PLACE (
   ID                   SERIAL               not null,
   PLA_ID               INT4                 null,
   NAME               VARCHAR(200)         null,
   constraint PK_PLACE primary key (ID)
);

/*==============================================================*/
/* Index: PLACE_PK                                              */
/*==============================================================*/
create unique index PLACE_PK on PLACE (
ID
);

/*==============================================================*/
/* Index: RELATIONSHIP_9_FK                                     */
/*==============================================================*/
create  index RELATIONSHIP_9_FK on PLACE (
PLA_ID
);

/*==============================================================*/
/* Table: SITE                                                  */
/*==============================================================*/
create table SITE (
   ID                   SERIAL               not null,
   CAT_ID               INT4                 null,
   PLA_ID               INT4                 null,
   NAME                 VARCHAR(200)         null,
   DESCRIPTION          VARCHAR(500)         null,
   LAT                  FLOAT8               null,
   LON                  FLOAT8               null,
   OFFICE_HOURS         INT4                 null,
   PHONE                VARCHAR(10)          null,
   WEB_SITE             VARCHAR(200)         null,
   THUMBNAIL            VARCHAR(200)         null,
   GEOM                 VARCHAR(200)         null,
   constraint PK_SITE primary key (ID)
);

/*==============================================================*/
/* Index: SITE_PK                                               */
/*==============================================================*/
create unique index SITE_PK on SITE (
ID
);

/*==============================================================*/
/* Index: RELATIONSHIP_8_FK                                     */
/*==============================================================*/
create  index RELATIONSHIP_8_FK on SITE (
PLA_ID
);

/*==============================================================*/
/* Index: RELATIONSHIP_11_FK                                    */
/*==============================================================*/
create  index RELATIONSHIP_11_FK on SITE (
CAT_ID
);

/*==============================================================*/
/* Table: USERS                                                 */
/*==============================================================*/

create table USERS (
   USERNAME             VARCHAR(100)               not null,
   Hashed_password       TEXT                 not null,
   FULL_NAME             VARCHAR(100)         null,
   EMAIL                VARCHAR(100)         null,
   password_changed_at     timestamp     null,
   created_at            timestamp           null,
   
   constraint PK_USERS primary key (USERNAME)
);

/*==============================================================*/
/* Index: USER_PK                                               */
/*==============================================================*/
create unique index USER_PK on USERS (
USERNAME
);

/*==============================================================*/
/* Table: VISIT                                                 */
/*==============================================================*/
create table VISIT (
   ID                   SERIAL               not null,
   SIT_ID               INT4                 not null,
   USERNAME             VARCHAR(100)          not null,
   DATE                 DATE                 null,
   constraint PK_VISIT primary key (ID)
);

/*==============================================================*/
/* Index: RELATIONSHIP_12_FK                                    */
/*==============================================================*/
create  index RELATIONSHIP_12_FK on VISIT (
SIT_ID
);

/*==============================================================*/
/* Index: RELATIONSHIP_13_FK                                    */
/*==============================================================*/
create  index RELATIONSHIP_13_FK on VISIT (
USERNAME
);

alter table ASSETS
   add constraint FK_ASSETS_RELATIONS_KIND foreign key (KIN_ID)
      references KIND (ID)
      on delete restrict on update restrict;

alter table ASSTES_SITE
   add constraint FK_ASSTES_S_RELATIONS_ASSETS foreign key (ASS_ID)
      references ASSETS (ID)
      on delete restrict on update restrict;

alter table ASSTES_SITE
   add constraint FK_ASSTES_S_RELATIONS_SITE foreign key (SIT_ID)
      references SITE (ID)
      on delete restrict on update restrict;

alter table COMMENTS
   add constraint FK_COMMENTS_RELATIONS_VISIT foreign key (VIS_ID)
      references VISIT (ID)
      on delete restrict on update restrict;

alter table PLACE
   add constraint FK_PLACE_RELATIONS_PLACE foreign key (PLA_ID)
      references PLACE (ID)
      on delete restrict on update restrict;

alter table SITE
   add constraint FK_SITE_RELATIONS_CATEGORY foreign key (CAT_ID)
      references CATEGORY (ID)
      on delete restrict on update restrict;

alter table SITE
   add constraint FK_SITE_RELATIONS_PLACE foreign key (PLA_ID)
      references PLACE (ID)
      on delete restrict on update restrict;

alter table VISIT
   add constraint FK_VISIT_RELATIONS_SITE foreign key (SIT_ID)
      references SITE (ID)
      on delete restrict on update restrict;

alter table VISIT
   add constraint FK_VISIT_RELATIONS_USERS foreign key (username)
      references USERS (username)
      on delete restrict on update restrict;

