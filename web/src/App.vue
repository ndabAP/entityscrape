<template>
  <div id="app">
    <veui-grid-container>
      <h1>Entity scrape</h1>
        <veui-grid-row v-for="entity in entities" :key="entity.entity">
          <veui-grid-column :span="24">
            <h2>{{ entity.entity }}</h2>
            <ul>
              <li v-for="adjective in entity.adjectives" :key="adjective.adjective">
                {{ adjective.adjective }}
              </li>
            </ul>
          </veui-grid-column>
        </veui-grid-row>

        <veui-grid-row>
          <veui-grid-column :span="24">
            <div id="mountNode"></div>
          </veui-grid-column>
        </veui-grid-row>
      </veui-grid-container>
  </div>
</template>

<script>
import { GridContainer, GridRow, GridColumn } from 'veui'

export default {
  components: {
    'veui-grid-container': GridContainer,
    'veui-grid-row': GridRow,
    'veui-grid-column': GridColumn
  },

  data: () => ({
    entities: [
      {
        entity: 'Donald John Trump',
        adjectives: [
          {
            adjective: 'cool',
            weight: 5
          },
          {
            adjective: 'nice',
            weight: 2
          },
          {
            adjective: 'sweet',
            weight: 9
          }
        ]
      },
      {
        entity: 'Barack Hussein Obama II',
        adjectives: [
          {
            adjective: 'great',
            weight: 2
          },
          {
            adjective: 'special',
            weight: 3
          },
          {
            adjective: 'dancing',
            weight: 7
          }
        ]
      }
    ]
  }),

  mounted () { }
}
</script>

<style lang="less">
@import "~veui-theme-one/common.less";

li {
  display:inline;
}
</style>
